// Package kMeans provides an implementation for the k-means algorithm with greedy initialization
package kMeans

import (
	openapi "backend/internal/swag_gen"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// A coordinates type contains longitude and latitude
type coordinates struct {
	latitude  float64
	longitude float64
}

// Type to store an observation (location of an activity)
type observation struct {
	id          int64
	coordinates coordinates
}

// Type that defines a cluster
type cluster struct {
	center       coordinates
	observations observations
}

// Type that defines multiple observations
type observations []observation

// Type that defines multiple clusters
type clusters []cluster

// Type that defines multiple activity models
type activities []openapi.ActivityModel

var ErrNrObservationsTooSmall = errors.New("number of activities smaller than number of clusters")
var ErrNrCandidateClustersTooSmall = errors.New("number of candidate clusters smaller than or equal to 0")
var ErrNoObservations = errors.New("number of candidates smaller than or equal to 0")
var ErrNoClusters = errors.New("number of clusters smaller than or equal to 0")

var zeroCoordinate = coordinates{
	latitude:  0,
	longitude: 0,
}

var zeroObservation = observation{
	id:          0,
	coordinates: zeroCoordinate,
}

var zeroObservations = observations{
	zeroObservation,
}

var zeroCluster = cluster{
	center:       zeroCoordinate,
	observations: zeroObservations,
}

var zeroClusters = clusters{
	zeroCluster,
}

func kMeans(activities activities, nrClusters int, nrCandidateClusters int) (clusters, error) {
	//variable used to store error and return when a function returns an error.
	var err error
	var observations = make(observations, 0)
	observations, err = activitiesToObservations(activities, observations)
	//return if error raises while converting activities to observations
	if err != nil {
		return zeroClusters, err
	}
	var clusters = make(clusters, 0, nrClusters)
	clusters, err = initializeClusters(observations, clusters, nrClusters, nrCandidateClusters)
	//return if error raises while initializing clusters
	if err != nil {
		return zeroClusters, err
	}

	converged := false

	//stop if converged
	for !converged {

		oldClusters := clusters
		//clear all the assigned activities
		clearObservations(&clusters)

		//for each activity calculate the closest cluster and assign activity to that cluster
		for _, observation := range observations {

			_, closestCenter, err := distanceToNearestCluster(observation, clusters)
			if err != nil {
				return zeroClusters, fmt.Errorf("got an error in distanceToNearestCluster: %v", err)
			}
			//assign observation to closest cluster
			clusters[closestCenter].observations = append(clusters[closestCenter].observations, observation)
		}

		//update the clusters
		clusters = updateCluster(clusters)

		if reflect.DeepEqual(clusters, oldClusters) {
			converged = true
		}

	}
	return clusters, err
}

// updateCluster requires a pointer to a clusters(list of cluster) and updates all centers of the given clusters.
func updateCluster(clusters clusters) clusters {
	for index, cluster := range clusters {
		sumLatitude := 0.00
		sumlongitude := 0.00
		//calculate the sum of all longitude / latitudes of points that are closest to the cluster
		for _, observation := range cluster.observations {
			sumLatitude += observation.coordinates.latitude
			sumlongitude += observation.coordinates.longitude
		}
		//Set new center equal to the average of all the points
		clusters[index].center = coordinates{
			latitude:  sumLatitude / float64(len(cluster.observations)),
			longitude: sumlongitude / float64(len(cluster.observations)),
		}
	}
	return clusters
}

// clearObservations requires a pointer to a clusters (list of cluster) and removes all observations from the cluster
func clearObservations(clusterPtr *clusters) {
	//derefence
	clusters := *clusterPtr
	for i := range clusters {
		clusters[i].observations = make(observations, 0)
	}
}

// initializeClusters requires a pointer to an observations (list of observation), a pointer to an clusters (list of cluster) and nrClusters
// the function will append nrClusters cluster to the clusters list. The centers of the clusters are chosen by the greedy k-means++ algorithm
// this algorithm is slower in intializing because it has to calculate the cost to the nearest center multiple times, but generally this way of initializing results
// can ensure that the cluster centers are well-spaced and representative of the data, reducing the chances of getting stuck in suboptimal local minima.
func initializeClusters(observations observations, clusters clusters, nrClusters int, nrCandidateClusters int) (clusters, error) {

	//throw error if nrCluster is larger than length of observations
	if nrClusters > len(observations) {
		return nil, ErrNrObservationsTooSmall
	}
	//throw error if nr of candidate clusters is smaller or equal to 0
	if nrCandidateClusters <= 0 {
		return nil, ErrNrCandidateClustersTooSmall
	}

	//randomly choose first cluster
	firstClusterIndex := rand.Intn(len(observations))
	firstObservation := make([]observation, 0)
	firstCluster := cluster{
		center:       observations[firstClusterIndex].coordinates,
		observations: append(firstObservation, observations[firstClusterIndex]),
	}
	clusters = append(clusters, firstCluster)

	// Select the next cluster centers one at a time, based on the distance from the existing centers

	// For each point calculate the distance to the nearest cluster center
	distances, _, err := totalListDistance(observations, clusters)
	if err != nil {
		return zeroClusters, fmt.Errorf("got an error in totalListDistance: %v", err)
	}

	for len(clusters) < nrClusters {

		probabilities := Probabilities(observations, distances)

		candidates := chooseCandidates(observations, probabilities, nrCandidateClusters, rand.New(rand.NewSource(time.Now().UnixNano())))

		// Choose the candidate with the minimum overall distance of all points to their closest center
		bestCandidate, err := bestCandidateFunc(candidates, clusters, observations)
		if err != nil {
			return clusters, fmt.Errorf("got error in call to bestCandidateFunc: %w", err)
		}
		// Append best candidate cluster to clusters
		clusters = append(clusters, bestCandidate)

	}
	// If finished without an error return nil
	return clusters, nil
}

func chooseCandidates(observations observations, probabilities []float64, nrCandidateClusters int, rand *rand.Rand) clusters {
	// Choose the next clusterCandidate centers randomly with probability proportional to the distance
	candidates := make([]cluster, 0, nrCandidateClusters)
	for i := 0; i < nrCandidateClusters; i++ {
		// Algo for running sum this could be used to improve efficiency when observations is really large this has to only loop over probabilties once
		// sum := 0.0
		// selectedIdx := -1
		// for i, p := range probs {
		// 	sum += p
		// 	if rand.Float64()*sum < p {
		// 		selectedIdx = i
		// 	}
		// }
		// If no data point was selected, return a default value
		// if selectedIdx == -1 {
		// 	return selectedIdx
		// }
		sum := 0.0
		for _, p := range probabilities {
			sum += p
		}
		r := rand.Float64() * sum
		for i, p := range probabilities {
			r -= p
			if r <= 0.0 {
				coordinates := coordinates{
					latitude:  observations[i].coordinates.latitude,
					longitude: observations[i].coordinates.longitude,
				}
				//cannot go out of bounds since probabilties is the same length as observations and index starts at 0
				observationChosen := observations[i : i+1]

				candidates = append(candidates, cluster{
					center:       coordinates,
					observations: observationChosen,
				})
				break

			}
		}
	}
	return candidates
}

// totalSumDistance takes observations (list of observation) and clusters (list of cluster) as input and returns a list that contains for each observation the distance to the closest cluster and
// a list that contains for each observation the index of the closest cluster
func totalSumDistance(observations observations, checkCluster clusters) (float64, error) {
	if len(checkCluster) <= 0 {
		return 0.0, ErrNoClusters
	}
	sumDistance := 0.0
	for _, observation := range observations {
		distance, _, err := distanceToNearestCluster(observation, checkCluster)
		sumDistance += distance
		//return if distanceToNearestCluster returned an error
		if err != nil {
			return 0, fmt.Errorf("got error in call to distanceToNearestCluster: %w", err)
		}
	}
	return sumDistance, nil
}

// totalListDistance takes observations (list of observation) and clusters (list of cluster) as input and returns a list that contains for each observation the distance to the closest cluster and
// a list that contains for each observation the index of the closest cluster
func totalListDistance(observations observations, clusters clusters) ([]float64, []int, error) {
	if len(clusters) <= 0 {
		return []float64{}, []int{}, ErrNoClusters
	}
	// For each point calculate the distance to the nearest cluster center
	distances := make([]float64, len(observations))
	closestCluster := make([]int, len(observations))
	for index, observation := range observations {
		var err error
		distances[index], closestCluster[index], err = distanceToNearestCluster(observation, clusters)
		if err != nil {
			return []float64{}, []int{}, fmt.Errorf("got an error in distanceToNearestCluster: %v", err)
		}
	}
	return distances, closestCluster, nil
}

func Probabilities(observations observations, distances []float64) []float64 {
	//calculate probabilities
	probabilities := make([]float64, len(observations))
	// Squaring the distance makes the probability of selecting a data point decrease more steeply as the distance increases, compared to using the distance itself.
	// This has the effect of encouraging the algorithm to select cluster centers that are farther apart and better representative of the data.
	for index, distance := range distances {
		probabilities[index] = distance * distance
	}
	return probabilities
	// Explanation of proporional probability because i failed probability and statistics and don't get it :(
	// To choose the next cluster center randomly with probability proportional to the distance,
	// generate a random number r between 0 and the sum of the probabilities, which is r = 38.9.
	// Then, iterate over the probabilities, subtracting each probability from r until r becomes negative.
	// The data point corresponding to the first negative value is selected as the next cluster center.

	// For example, suppose that we subtract the probabilities in order from r and reach a negative value after subtracting the sixth probability

	// r = 38.9 - 5.29 = 33.61
	// r = 33.61 - 2.25 = 31.36
	// r = 31.36 - 0.64 = 30.72
	// r = 30.72 - 10.24 = 20.48
	// r = 20.48 - 1.21 = 19.27
	// r = 19.27 - 8.41 = 10.86 (negative)
	// In this case, the data point corresponding to the sixth probability (with distance 2.9) is selected as the next cluster center.
}

//func selectCandidate()

// bestCandidateFunc requires as input a candidates (list of cluster), clusters and observations. Given these 3 the function calculates and returns
// which candidate minimizes the total distance when added to the clusters
func bestCandidateFunc(candidates clusters, clusters clusters, observations observations) (cluster, error) {

	if len(candidates) <= 0 {
		return zeroCluster, ErrNrCandidateClustersTooSmall
	}
	bestCandidate := candidates[0]
	bestDistance := math.Inf(1)
	// Loop over all candidate centers
	for index, candidate := range candidates {
		//candidate to clusters
		checkCluster := append(clusters, candidate)

		//calculate the sum of distances and if smaller this is the best candidate so far
		sumDistance, err := totalSumDistance(observations, checkCluster)
		if err != nil {
			return cluster{}, fmt.Errorf("got error in call to totalSumDistance: %w", err)
		}

		//if the sumDistance of this candidate is lower replace the bestDistance and bestCandidate
		if sumDistance < bestDistance {
			bestDistance = sumDistance
			bestCandidate = candidates[index]
		}
	}
	return bestCandidate, nil
}

// distanceToNearestCluster requires a pointer to an observations (list of observation) and a pointer to an clusters (list of cluster)
// and returns the cluster that is nearest and the index to the cluster
func distanceToNearestCluster(observation observation, clusters clusters) (float64, int, error) {
	//throw error if nr of observations is smaller than or equal to 0
	if clusters == nil {
		return 0, 0, errors.New("no observation passed in distanceToNearestCluster")
	}

	// Compute the distance from x to the nearest cluster in clusters
	minDist := math.Inf(1)
	minCluster := 0
	for index, cluster := range clusters {
		dist := distance(observation, cluster)
		if dist < minDist {
			minDist = dist
			minCluster = index
		}
	}
	return minDist, minCluster, nil
}

// distance takes as input an observation and a cluster and calculates the difference between the observation and the cluster
func distance(observation observation, cluster cluster) float64 {
	firstArgument := observation.coordinates.latitude - cluster.center.latitude
	secondArgument := observation.coordinates.longitude - cluster.center.longitude
	return math.Sqrt((firstArgument * firstArgument) + (secondArgument * secondArgument))
}

// activitiesToObservations requires a pointer to an activities (list of ActivityModel) and a pointer to an observations (list of observation)
// converts the activity.AddressApplied latitude and longitude to an observation langtitude and longitude and grabs the activity.Id, combines these two into an
// observation and appends this to the observations.
func activitiesToObservations(activities activities, observations observations) (observations, error) {

	//throw error if nr of observations is smaller than or equal to 0
	if len(activities) <= 0 {
		return nil, ErrNoObservations
	}

	var newId int64
	//loop over all activities in activities (activities is a list of openapi.ActivityModel)
	for _, activity := range activities {
		newId = activity.GetId()

		newLatitude, err := strconv.ParseFloat(*activity.AddressApplied.Latitude, 64)
		//return if error raises while parsing
		if err != nil {
			return nil, err
		}
		newlongitude, err := strconv.ParseFloat(*activity.AddressApplied.Longitude, 64)
		//return if error raises while parsing
		if err != nil {
			return nil, err
		}
		//create new coordinates using the long and lat we got from the activity
		newCoordinates := coordinates{
			latitude:  newLatitude,
			longitude: newlongitude,
		}

		//create new observation with newCoordinates and id of the activity
		newObservation := observation{
			id:          newId,
			coordinates: newCoordinates,
		}
		observations = append(observations, newObservation)
	}
	return observations, nil
}
