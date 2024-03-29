// Package kMeans provides an implementation for the k-means algorithm with greedy initialization
package kMeans

import (
	model "bzone/backend/internal/models"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// A Coordinates type contains Longitude and Latitude
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Type to store an observation (location of an activity)
type observation struct {
	id          int64
	Coordinates Coordinates
}

// Type that defines a Cluster
type Cluster struct {
	Center       Coordinates
	observations observations
}

// Type that defines multiple observations
type observations []observation

// Type that defines multiple Clusters
type Clusters []Cluster

// Type that defines multiple activity models
type activities []model.ActivityModelBumbal

// radius of the Earth in kilometers (for haversine distance)
const earthRadius = 6371

// Standard errors that are returned
var ErrNrObservationsTooSmall = errors.New("number of activities smaller than number of Clusters")
var ErrNrCandidateClustersTooSmall = errors.New("number of candidate Clusters smaller than or equal to 0")
var ErrNoObservations = errors.New("number of candidates smaller than or equal to 0")
var ErrNoClusters = errors.New("number of Clusters smaller than or equal to 0")
var ErrNoActivities = errors.New("number of activities smaller than or equal to 0")
var ErrNrClustersTooSmall = errors.New("nrClusters variable is smaller than or equal to 0")

// Standard empty structs to return when an error occurs
var zeroCoordinate = Coordinates{
	Latitude:  0,
	Longitude: 0,
}

var zeroObservation = observation{
	id:          0,
	Coordinates: zeroCoordinate,
}

var zeroObservations = observations{
	zeroObservation,
}

var zeroCluster = Cluster{
	Center:       zeroCoordinate,
	observations: zeroObservations,
}

// KMeans is an implementation of the k-means clustering algorithm. Given a set of activities, it uses the k-means algorithm to group the activities into a specified number of clusters.
// The function takes in the activities, the desired number of clusters, and the desired number of candidate clusters to use during initialization. If any of the input parameters do not meet the preconditions, an error is returned.
// The function initializes the clusters with random centers and then repeatedly assigns observations to their nearest cluster and updates the cluster centers until convergence is achieved.
// The function returns the resulting clusters and any errors that were encountered during the clustering process.
func KMeans(activities activities, nrClusters int, nrCandidateClusters int) (Clusters, error) {
	// return error if precondition isn't met
	if len(activities) <= 0 {
		return nil, ErrNoActivities
	}

	if nrClusters <= 0 {
		return nil, ErrNrClustersTooSmall
	}

	//variable to store observations
	var observations = make(observations, 0)

	observations, err := activitiesToObservations(activities, observations)

	//return if error raises while converting activities to observations
	if err != nil {
		return nil, err
	}

	//variable to store Clusters
	var Clusters = make(Clusters, 0, nrClusters)

	//initialize Clusters with random source of current time
	Clusters, err = initializeClusters(observations, Clusters, nrClusters, nrCandidateClusters, rand.New(rand.NewSource(time.Now().UnixNano())))

	//return if error raises while initializing Clusters
	if err != nil {
		return nil, err
	}

	converged := false

	//stop if converged
	for !converged {

		oldClusters := Clusters
		//clear all the assigned activities
		clearObservations(&Clusters)

		//for each observation calculate the closest Cluster and assign observation to that Cluster
		for _, observation := range observations {

			_, closestCenter, err := distanceToNearestCluster(observation, Clusters)

			//return if error raises in distanceToNearestCluster
			if err != nil {
				return nil, fmt.Errorf("got an error in distanceToNearestCluster: %v", err)
			}

			//assign observation to the closest Cluster
			Clusters[closestCenter].observations = append(Clusters[closestCenter].observations, observation)
		}

		//update the Clusters
		Clusters := updateCluster(Clusters)

		//stop if oldClusters are equal to updated Clusters
		if reflect.DeepEqual(Clusters, oldClusters) {
			converged = true
		}

	}

	return Clusters, err
}

// updateCluster requires Clusters, for each Cluster in Clusters it will calculate the center of all associated observations
// and set the center field of the Cluster to the calculated center. Returns the updates Clusters.
func updateCluster(Clusters Clusters) Clusters {
	for index, Cluster := range Clusters {
		if len(Cluster.observations) == 0 {
			// skip over empty Clusters
			continue
		}

		// initiate sum variables
		sumLatitude := 0.00
		sumlongitude := 0.00

		// calculate the sum of all Longitude / Latitude of points assigned to the Cluster
		for _, observation := range Cluster.observations {
			sumLatitude += observation.Coordinates.Latitude
			sumlongitude += observation.Coordinates.Longitude
		}

		// set new Center equal to the average of all the points assigned to the Cluster
		newCenter := Coordinates{
			Latitude:  sumLatitude / float64(len(Cluster.observations)),
			Longitude: sumlongitude / float64(len(Cluster.observations)),
		}

		// update Center of the current Cluster
		Clusters[index].Center = newCenter
	}

	return Clusters
}

// clearObservations requires a pointer to a Clusters (list of Cluster) and removes all observations from the Cluster
func clearObservations(clusterPtr *Clusters) {
	//derefence
	Clusters := *clusterPtr
	for i := range Clusters {
		Clusters[i].observations = make(observations, 0)
	}
}

// initializeClusters requires a pointer to an observations (list of observation), a pointer to an Clusters (list of Cluster) and nrClusters
// the function will append nrClusters Cluster to the Clusters list. The centers of the Clusters are chosen by the greedy k-means++ algorithm
// this algorithm is slower in intializing because it has to calculate the cost to the nearest Center multiple times, but generally this way of initializing results
// can ensure that the Cluster centers are well-spaced and representative of the data, reducing the chances of getting stuck in suboptimal local minima.
func initializeClusters(observations observations, Clusters Clusters, nrClusters int, nrCandidateClusters int, rand *rand.Rand) (Clusters, error) {

	//throw error if nrCluster is larger than length of observations
	if nrClusters > len(observations) {
		return nil, ErrNrObservationsTooSmall
	}
	//throw error if nr of candidate Clusters is smaller or equal to 0
	if nrCandidateClusters <= 0 {
		return nil, ErrNrCandidateClustersTooSmall
	}
	if nrCandidateClusters > len(observations) {
		return nil, ErrNrObservationsTooSmall
	}

	//randomly choose first Cluster
	firstClusterIndex := rand.Intn(len(observations))
	firstObservation := make([]observation, 0)
	firstCluster := Cluster{
		Center:       observations[firstClusterIndex].Coordinates,
		observations: append(firstObservation, observations[firstClusterIndex]),
	}
	Clusters = append(Clusters, firstCluster)

	// Select the next Cluster centers one at a time, based on the distance from the existing centers

	// For each point calculate the distance to the last chosen candidate Center
	distances := totalListDistance(observations, Clusters[len(Clusters)-1])

	for len(Clusters) < nrClusters {

		probabilities := Probabilities(distances)

		candidates := chooseCandidates(observations, probabilities, nrCandidateClusters, rand)

		// Choose the candidate with the minimum overall distance of all points to their closest Center
		bestCandidate, err := bestCandidateFunc(candidates, Clusters, observations)
		if err != nil {
			return Clusters, fmt.Errorf("got error in call to bestCandidateFunc: %w", err)
		}
		// Append the best candidate Cluster to Clusters
		Clusters = append(Clusters, bestCandidate)

	}
	// If finished without an error return nil
	return Clusters, nil
}

// chooseCandidates selects candidate Cluster centers randomly with probability proportional to the distance to the nearest existing Center.
func chooseCandidates(observations observations, probabilities []float64, nrCandidateClusters int, r *rand.Rand) Clusters {
	candidates := make([]Cluster, 0, nrCandidateClusters)

	// Select nrCandidateClusters number of candidates
	for i := 0; i < nrCandidateClusters; i++ {
		// Algo for running sum could be used to improve efficiency when observations is really large that has to loop over probabilties only once
		// Calculate the sum of all probabilities
		sum := 0.0
		for _, p := range probabilities {
			sum += p
		}
		// Generate a random number between 0 and the sum of all probabilities
		randomNumber := r.Float64() * sum
		// Select the observation that corresponds to the first probability value that is greater than or equal to the random number
		for i, p := range probabilities {
			randomNumber -= p
			if randomNumber <= 0.0 {
				// Create a new Cluster centered at the selected observation
				Coordinates := Coordinates{
					Latitude:  observations[i].Coordinates.Latitude,
					Longitude: observations[i].Coordinates.Longitude,
				}
				observationChosen := observations[i : i+1]

				candidates = append(candidates, Cluster{
					Center:       Coordinates,
					observations: observationChosen,
				})
				break
			}
		}
	}

	return candidates
}

// totalSumDistance calculates the total sum of distances of each observation to its closest Cluster and returns it.
// It also returns an error if the provided list of Clusters is empty.
func totalSumDistance(observations observations, checkCluster Clusters) (float64, error) {
	// Return an error if the list of Clusters is empty
	if len(checkCluster) <= 0 {
		return 0.0, ErrNoClusters
	}

	sumDistance := 0.0
	// Calculate the distance of each observation to its closest Cluster and add it to the sumDistance variable
	for _, observation := range observations {
		distance, _, err := distanceToNearestCluster(observation, checkCluster)
		// If distanceToNearestCluster returned an error, return it as an error from this function
		if err != nil {
			return 0, fmt.Errorf("got error in call to distanceToNearestCluster: %w", err)
		}
		sumDistance += distance
	}

	// Return the total sum of distances to the closest Clusters
	return sumDistance, nil
}

// totalListDistance calculates the distances of each observation in the provided observations slice to the given Cluster Center and returns a slice of distances.
func totalListDistance(observations observations, Cluster Cluster) []float64 {
	// For each observation, calculate the distance to the provided Cluster Center
	distances := make([]float64, len(observations))
	for index, observation := range observations {
		// Calculate the distance between the observation and the Cluster Center
		distances[index] = distance(observation, Cluster)

	}
	// Return the slice of distances
	return distances
}

// Probabilities calculates the probability of selecting each observation from the provided distances to be proportional to the squared distance.
// It returns a slice of probabilities for each observation.
func Probabilities(distances []float64) []float64 {
	// Initialize a slice to store the probabilities
	probabilities := make([]float64, len(distances))

	// Calculate the probability of selecting each observation proportional to the squared distance
	for index, distance := range distances {
		// Squaring the distance makes the probability of selecting a data point decrease more steeply as the distance increases, compared to using the distance itself.
		// This has the effect of encouraging the algorithm to select Cluster centers that are farther apart and better representative of the data.
		probabilities[index] = distance * distance
	}

	// Return the slice of probabilities
	return probabilities
}

// bestCandidateFunc takes a list of candidate Cluster centers, a list of existing Clusters, and a list of observations as input.
// It calculates the total sum of distances of each observation to its closest Cluster when each candidate Center is added to the existing Clusters,
// and returns the candidate Center that minimizes the total distance, along with no error. It returns an error if there is an error in the distance calculation or if the list of candidates is empty.
func bestCandidateFunc(candidates Clusters, Clusters Clusters, observations observations) (Cluster, error) {
	// Return an error if the list of candidates is empty
	if len(candidates) <= 0 {
		return zeroCluster, ErrNrCandidateClustersTooSmall
	}

	// Initialize the best candidate and best distance variables
	bestCandidate := candidates[0]
	bestDistance := math.Inf(1)

	// Loop over all candidate centers
	for index, candidate := range candidates {
		// Add the candidate Center to the existing Clusters
		checkCluster := append(Clusters, candidate)

		// Calculate the total sum of distances of each observation to its closest Cluster when the candidate Center is added to the existing Clusters
		sumDistance, err := totalSumDistance(observations, checkCluster)
		if err != nil {
			return Cluster{}, fmt.Errorf("got error in call to totalSumDistance: %w", err)
		}

		// If the sumDistance of this candidate is lower than the current bestDistance, replace the bestDistance and bestCandidate variables
		if sumDistance < bestDistance {
			bestDistance = sumDistance
			bestCandidate = candidates[index]
		}
	}

	// Return the candidate Center that minimizes the total distance and no error
	return bestCandidate, nil
}

// distanceToNearestCluster takes an observation and a list of Clusters as input, and calculates the distance from the observation to the nearest Cluster.
// It returns the distance to the nearest Cluster, the index of the nearest Cluster, and no error.
// It returns an error if the list of Clusters is nil.
func distanceToNearestCluster(observation observation, Clusters Clusters) (float64, int, error) {
	// Throw an error if the list of Clusters is nil
	if Clusters == nil {
		return 0, 0, errors.New("no observation passed in distanceToNearestCluster")
	}

	// Initialize the minimum distance and minimum Cluster variables
	minDist := math.Inf(1)
	minCluster := 0

	// Loop over all Clusters and calculate the distance to the observation
	for index, Cluster := range Clusters {
		dist := distance(observation, Cluster)

		// If the distance to this Cluster is smaller than the current minimum distance, update the minimum distance and minimum Cluster variables
		if dist < minDist {
			minDist = dist
			minCluster = index
		}
	}

	// Return the minimum distance and minimum Cluster index with no error
	return minDist, minCluster, nil
}

// distance takes as input an observation and a Cluster and calculates the difference between the observation and the Cluster
func distance(observation observation, Cluster Cluster) float64 {
	// Calculate the difference between the observation and the Cluster for both the Latitude and Longitude
	latitudeDifference := observation.Coordinates.Latitude - Cluster.Center.Latitude
	longitudeDifference := observation.Coordinates.Longitude - Cluster.Center.Longitude

	// Calculate the distance between the observation and the Cluster using the Pythagorean theorem
	distance := math.Sqrt((latitudeDifference * latitudeDifference) + (longitudeDifference * longitudeDifference))

	// Return the distance
	return distance
}

// distanceKilometers takes as input an observation and a Cluster and calculates the distance between the observation and the Cluster in kilometers
// distanceKilometers can replace distance for a more accurate result (taking into account the spherical nature of earth)
func distanceKilometers(observation observation, Cluster Cluster) float64 {
	// Calculate the distance between the observation and the Cluster using the haversine formula
	resultHaversine := haversineDistance(observation.Coordinates, Cluster.Center)

	// Return the distance in kilometers
	return resultHaversine
}

// activitiesToObservations takes a list of activities and a list of observations as input, and converts the Latitude and Longitude values from the activities to observations.
// It appends each new observation to the observations list and returns the updated list.
// It throws an error if the number of activities is smaller than or equal to 0.
func activitiesToObservations(activities activities, observations observations) (observations, error) {
	// Throw an error if the number of activities is smaller than or equal to 0
	if len(activities) <= 0 {
		return nil, ErrNoObservations
	}

	// Loop over all activities in activities (activities is a list of bumbal.ActivityModel)
	for _, activity := range activities {
		// Convert the activity ID to an integer
		newId, err := strconv.ParseInt(activity.GetId(), 10, 64)
		// Return an error if an error occurs while parsing the ID
		if err != nil {
			return nil, err
		}

		// Convert the activity Latitude to a float
		newLatitude, err := strconv.ParseFloat(*activity.AddressApplied.Latitude, 64)
		// Return an error if an error occurs while parsing the Latitude
		if err != nil {
			return nil, err
		}

		// Convert the activity Longitude to a float
		newLongitude, err := strconv.ParseFloat(*activity.AddressApplied.Longitude, 64)
		// Return an error if an error occurs while parsing the Longitude
		if err != nil {
			return nil, err
		}

		// Create a new Coordinates struct using the Latitude and Longitude values
		newCoordinates := Coordinates{
			Latitude:  newLatitude,
			Longitude: newLongitude,
		}

		// Create a new observation with the new Coordinates and the activity ID
		newObservation := observation{
			id:          newId,
			Coordinates: newCoordinates,
		}

		// Append the new observation to the observations list
		observations = append(observations, newObservation)
	}

	// Return the updated observations list and no error
	return observations, nil
}

// toRadians takes a float64 value in degrees and converts it to radians
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// haversineDistance takes two Coordinates as input and calculates the distance between them using the Haversine formula
func haversineDistance(point1, point2 Coordinates) float64 {
	// Convert Latitude and Longitude to radians
	lat1 := toRadians(point1.Latitude)
	lon1 := toRadians(point1.Longitude)
	lat2 := toRadians(point2.Latitude)
	lon2 := toRadians(point2.Longitude)

	// Calculate differences in Latitude and Longitude
	deltaLat := lat2 - lat1
	deltaLon := lon2 - lon1

	// Calculate the central angle between the two points
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance using the Haversine formula and the Earth's radius
	distance := earthRadius * c

	return distance
}
