package bumbal

import (
	"bytes"
	"bzone/backend/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const baseUrl = "https://sep202302.bumbal.eu/api/v2/"
const (
	zonesUrl = "zone"
)

func (b *BumbalZoneListResponse) ToRangeModelList() ([]models.ZoneModel, error) {

	zoneList := make([]models.ZoneModel, 0, len(b.Items))

	for _, bumbalZone := range b.Items {
		// for each bumbal zone create a Zone Model
		var newZone models.ZoneModel

		// for each zone range in the bumbal zone
		for _, bumbalZoneRange := range *bumbalZone.ZoneRanges {

			// create a new models.ZoneRangeModel
			// we only care about the zipcodes and the country
			var newZoneRange models.ZoneRangeModel
			z, err := strconv.Atoi(bumbalZoneRange.ZipcodeFrom)
			if err != nil {
				return nil, err
			}

			newZoneRange.ZipcodeFrom = z

			z, err = strconv.Atoi(bumbalZoneRange.ZipcodeTo)
			if err != nil {
				return nil, err
			}

			newZoneRange.ZipcodeTo = z
			newZoneRange.IsoCountry = bumbalZoneRange.IsoCountry

			// add the new zone range to the zone
			newZone.ZoneRanges = append(newZone.ZoneRanges, newZoneRange)
		}

		zoneList = append(zoneList, newZone)
	}

	return zoneList, nil
}

// helper function to convert a BumbalZoneListResponse received from bumbal's API to a PlotModel
// makes use of the ToRangeModelList function internally
func (b *BumbalZoneListResponse) ToPlotModel() (*models.PlotModel, error) {
	zones, err := b.ToRangeModelList()
	if err != nil {
		return nil, err
	}

	var plotModel models.PlotModel
	// I think this makes a copy of the slice, might want to pass a pointer? idk
	plotModel.PlotId = uuid.NewString()
	plotModel.Zones = zones
	// mark the origin of the zone as coming from bumbal
	plotModel.Origin = models.OriginBumbal
	plotModel.PlotCreatedAt = time.Now()
	plotModel.PlotSavedAt = time.Now()

	return &plotModel, nil
}

// GetZoneListReponse gets the zones from bumbal using the jwt token
// IMPORTANT: the jwt token must be valid
func GetZoneListReponse(jwt string) (*BumbalZoneListResponse, error) {

	// fail safe... I shot myself in the foot wit this one...
	// update: I shot myself in the foot again, I was using TrimLeft ...
	jwt = strings.TrimPrefix(jwt, "Bearer ")

	if jwt == "" {
		return nil, fmt.Errorf("jwt token must not be empty")
	}

	// get the zones from bumbal
	requestBody := []byte(`{
	  "options": {
	   "include_zone_ranges": true,
	   "include_brands": false
	  }
	}`)

	url := baseUrl + zonesUrl

	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("The JWT is not valid")
	}

	var resModel BumbalZoneListResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&resModel)
	if err != nil {
		return nil, err
	}

	fmt.Printf("res model %+v\n\n", resModel)

	return &resModel, nil
}
