package kMeans

import (
	model "bzone/backend/internal/models"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// clusterToZoneModel takes a list of Clusters and a list of activities as input and converts the Clusters to ZoneModels.
// It first converts the Clusters to sets of zipcodes using the clusterToZipcodeSet function.
// It then converts the sets of zipcodes to ZoneModels using the zipcodeSetToZoneModel function.
// If an error occurs in either of these functions, it returns the error.
// Otherwise, it returns a list of ZoneModels.
func ClusterToZoneModel(Clusters Clusters, activities activities) ([]model.ZoneModel, error) {
	zipcodeSets, err := clusterToZipcodeSet(Clusters, activities)
	if err != nil {
		return nil, err
	}
	zones, err := zipcodeSetToZoneModel(zipcodeSets)
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// clusterToZipcodeSet takes a list of Clusters and a list of activities as input and converts the Clusters to sets of zipcodes.
// It loops over each Cluster, and for each observation in the Cluster, it matches the observation ID to an activity ID to get the corresponding zipcode.
// It then adds the zipcode to a set for the Cluster.
// It returns a list of sets of zipcodes, one set per Cluster.
// If an error occurs while parsing the activity ID or getting the zipcode, it returns the error.
func clusterToZipcodeSet(Clusters Clusters, activities activities) ([]map[string]struct{}, error) {
	if len(activities) <= 0 {
		return nil, errors.New("length of activities is too small (clusterToZipCodeSet)")
	}

	// Make a slice that holds sets of zipcodes
	listZipcodeSet := make([]map[string]struct{}, 0)

	// Loop over each Cluster
	for _, Cluster := range Clusters {
		// Set that holds zipcodes for each Cluster
		zipcodeSet := make(map[string]struct{})

		// Loop over all observations in the Cluster and add the corresponding zipcode to the set
		for _, observation := range Cluster.observations {
			for _, activity := range activities {
				// Match the activity ID to the observation ID to get the corresponding zipcode
				activityId, err := strconv.ParseInt(activity.GetId(), 10, 64)
				if err != nil {
					return nil, err
				}
				if activityId == observation.id {
					zipcode := *activity.AddressApplied.Zipcode
					zipcodeSet[zipcode] = struct{}{}
				}
			}
		}

		// Append the set of zipcodes to the list of sets
		listZipcodeSet = append(listZipcodeSet, zipcodeSet)
	}

	// Return the list of sets of zipcodes and no error
	return listZipcodeSet, nil
}

// zipcodeSetToZoneModel takes a list of sets of zipcodes as input and converts each set to a ZoneModel.
// It loops over each set of zipcodes, and for each set, it calls the createZoneRanges function to create a list of ZoneRanges.
// It then creates a ZoneModel with the ZoneRanges, and adds it to a list of ZoneModels.
// The function returns the list of ZoneModels and no error.
// If an error occurs while creating the ZoneRanges, it returns the error.
func zipcodeSetToZoneModel(listZipcodeSet []map[string]struct{}) ([]model.ZoneModel, error) {
	idCounter := 0
	zones := make([]model.ZoneModel, 0)

	// Loop over each set of zipcodes
	for _, zipcodeSet := range listZipcodeSet {
		// If the set is empty, skip to the next set
		if len(zipcodeSet) == 0 {
			continue
		}

		// Call the createZoneRanges function to create a list of ZoneRanges for the set
		zoneRanges, err := createZoneRanges(idCounter, zipcodeSet)
		if err != nil {
			return nil, err
		}

		// Create a ZoneModel with the ZoneRanges and add it to the list of ZoneModels
		zone := model.ZoneModel{
			Id:              strconv.Itoa(idCounter),
			ZoneRanges:      zoneRanges,
			ZoneFuelCost:    0,
			ZoneDrivingTime: 0,
		}
		idCounter++
		zones = append(zones, zone)
	}

	// Return the list of ZoneModels and no error
	return zones, nil
}

// createZoneRanges takes an ID counter and a set of zipcodes as input and converts the set to a list of ZoneRanges.
// It first converts the zipcodes to integers and sorts them.
// It then loops over the sorted zipcodes and creates a new ZoneRange for each contiguous range of zipcodes.
// It adds each new ZoneRange to a list of ZoneRanges.
// The function then merges overlapping ZoneRanges and returns the merged list of ZoneRanges and no error.
// If an error occurs while creating or merging the ZoneRanges, it returns the error.
func createZoneRanges(idCounter int, zipcodeSet map[string]struct{}) ([]model.ZoneRangeModel, error) {
	if len(zipcodeSet) <= 0 {
		return nil, fmt.Errorf("length of zipcodeSet is too small: %q", zipcodeSet)
	}

	// Convert the zip codes to integers and sort them
	zipcodes := make([]int64, 0, len(zipcodeSet))
	for zipCode := range zipcodeSet {
		trimmedZipCode := strings.TrimRightFunc(zipCode, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		if len(trimmedZipCode) != 4 {
			return nil, fmt.Errorf("zipcode %q trimmed to %q is not 4 digits long", zipCode, trimmedZipCode)
		}
		zipcodeInt, err := strconv.ParseInt(trimmedZipCode, 10, 64)
		if err != nil {
			return nil, err
		}
		zipcodes = append(zipcodes, zipcodeInt)
	}
	sort.Slice(zipcodes, func(i, j int) bool {
		return zipcodes[i] < zipcodes[j]
	})

	// Create a list of ZoneRanges
	zoneRanges := make([]model.ZoneRangeModel, 0)

	var initialZipcode, lastZipcode int64
	firstLoop := true

	// Loop over each sorted zipcode and create a new ZoneRange for each contiguous range
	for _, zipcodeInt := range zipcodes {
		if firstLoop {
			// Initialize the initialZipcode and lastZipcode for the first range
			initialZipcode = zipcodeInt
			lastZipcode = zipcodeInt
			firstLoop = false
		} else {
			if lastZipcode+1 == zipcodeInt {
				// Expand the last range if the current zipcode is contiguous
				lastZipcode = zipcodeInt
			} else {
				// Create a new ZoneRangeModel for the previous range
				zoneRange, err := createZoneRange(strconv.Itoa(idCounter), initialZipcode, lastZipcode, "NLD")
				if err != nil {
					return nil, err
				}
				idCounter++
				zoneRanges = append(zoneRanges, zoneRange)
				// Update initialZipcode and lastZipcode for the next range
				initialZipcode = zipcodeInt
				lastZipcode = zipcodeInt
			}
		}
	}

	// Append the last zone range
	zoneRange, err := createZoneRange(strconv.Itoa(idCounter), initialZipcode, lastZipcode, "NLD")
	if err != nil {
		return nil, err
	}
	zoneRanges = append(zoneRanges, zoneRange)

	// Merge overlapping zone ranges
	mergedRanges := make([]model.ZoneRangeModel, 0, len(zoneRanges))

	for i := 0; i < len(zoneRanges); i++ {
		mergedRange := zoneRanges[i]

		for j := i + 1; j < len(zoneRanges); j++ {
			if zoneRanges[j].ZipcodeFrom <= mergedRange.ZipcodeTo+1 {
				mergedRange.ZipcodeTo = zoneRanges[j].ZipcodeTo
				i++
			} else {
				break
			}
		}

		mergedRanges = append(mergedRanges, mergedRange)
	}

	// Return the merged list of ZoneRanges and no error
	return mergedRanges, nil
}

// createZoneRange takes an ID string, a starting zipcode, an ending zipcode, and an ISO country code as input and returns a new ZoneRangeModel.
// It checks that the starting zipcode is not greater than the ending zipcode. If it is, the function returns an error.
// It then converts the ID string to an integer and initializes a new ZoneRangeModel with the converted ID, the starting and ending zipcodes, and the ISO country code.
// The function returns the new ZoneRangeModel and no error.
func createZoneRange(id string, zipcodeFrom int64, zipcodeTo int64, IsoCountry string) (model.ZoneRangeModel, error) {
	if zipcodeFrom > zipcodeTo {
		return model.ZoneRangeModel{}, fmt.Errorf("invalid zipcode range: zipcodeFrom (%d) cannot be greater than zipcodeTo (%d)", zipcodeFrom, zipcodeTo)
	}

	idInt, _ := strconv.ParseInt(id, 10, 64)
	//if err != nil {
	//	return nil, err
	//}

	// Create a new ZoneRangeModel with the given ID, zipcodes, and ISO country code
	zoneRange := model.ZoneRangeModel{
		ZoneRangeId: int(idInt),
		ZipcodeFrom: int(zipcodeFrom),
		ZipcodeTo:   int(zipcodeTo),
		IsoCountry:  IsoCountry,
	}

	// Return the new ZoneRangeModel and no error
	return zoneRange, nil
}
