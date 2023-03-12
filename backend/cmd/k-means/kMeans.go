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
)

// A coordinates type contains longtitude and latitude
type coordinates struct {
	longtitude float64
	latitude   float64
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
var ErrNrCandidateClustersTooSmall = errors.New("number of candidate smaller than or equal to 0")
var ErrNoObservations = errors.New("number of candidates smaller than or equal to 0")

var zeroCoordinate = coordinates{
	latitude:   0,
	longtitude: 0,
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

			_, closestCenter := distanceToNearestCluster(observation, clusters)
			//assign observation to closest cluster
			fmt.Println(len(clusters))

			fmt.Println(clusters)
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
		sumLongtitude := 0.00
		//calculate the sum of all longtitude / latitudes of points that are closest to the cluster
		for _, observation := range cluster.observations {
			sumLatitude += observation.coordinates.latitude
			sumLongtitude += observation.coordinates.longtitude
		}
		//Set new center equal to the average of all the points
		clusters[index].center = coordinates{
			latitude:   sumLatitude / float64(len(cluster.observations)),
			longtitude: sumLongtitude / float64(len(cluster.observations)),
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

	//throw error if nr of observations is smaller than or equal to 0
	if len(observations) <= 0 {
		return nil, ErrNoObservations
	}
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
	distances := make([]float64, len(observations))
	closestCluster := make([]int, len(observations))
	for index, observation := range observations {
		distances[index], closestCluster[index] = distanceToNearestCluster(observation, clusters)
	}

	for len(clusters) < nrClusters {

		//calculate probabilities
		probabilities := make([]float64, len(observations))
		// Squaring the distance makes the probability of selecting a data point decrease more steeply as the distance increases, compared to using the distance itself.
		// This has the effect of encouraging the algorithm to select cluster centers that are farther apart and better representative of the data.
		for index, distance := range distances {
			probabilities[index] = distance * distance
		}

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

		// Choose the next clusterCandidate centers randomly with probability proportional to the distance
		candidates := make([]cluster, nrCandidateClusters)
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
						latitude:   observations[i].coordinates.latitude,
						longtitude: observations[i].coordinates.longtitude,
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

		// Choose the candidate with the minimum overall distance of all points to their closest center
		bestCandidate := bestCandidateFunc(candidates, clusters, observations)
		// Append best candidate cluster to clusters
		clusters = append(clusters, bestCandidate)

	}
	// If finished without an error return nil
	return clusters, nil
}

func bestCandidateFunc(candidates clusters, clusters clusters, observations observations) cluster {

	bestCandidate := candidates[0]
	bestDistance := math.Inf(1)
	// Loop over all candidate centers
	for index, candidate := range candidates {
		checkCluster := append(clusters, candidate)
		sumDistance := 0.00
		//calculate the sum of distances and if smaller this is the best candidate so far
		for _, observation := range observations {
			distance, _ := distanceToNearestCluster(observation, checkCluster)
			sumDistance += distance
		}
		if sumDistance < bestDistance {
			bestDistance = sumDistance
			bestCandidate = candidates[index]
		}
	}
	return bestCandidate
}

// distanceToNearestCluster requires a pointer to an observations (list of observation) and a pointer to an clusters (list of cluster)
// and returns the cluster that is nearest and the index to the cluster
func distanceToNearestCluster(observation observation, clusters clusters) (float64, int) {
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
	return minDist, minCluster
}

// distance takes as input an observation and a cluster and calculates the difference between the observation and the cluster
func distance(observation observation, cluster cluster) float64 {
	firstArgument := observation.coordinates.latitude - cluster.center.latitude
	secondArgument := observation.coordinates.longtitude - cluster.center.longtitude
	return math.Sqrt((firstArgument * firstArgument) + (secondArgument * secondArgument))
}

// activitiesToObservations requires a pointer to an activities (list of ActivityModel) and a pointer to an observations (list of observation)
// converts the activity.AddressApplied latitude and longtitude to an observation langtitude and longtitude and grabs the activity.Id, combines these two into an
// observation and appends this to the observations.
func activitiesToObservations(activities activities, observations observations) (observations, error) {
	var newId int64
	//loop over all activities in activities (activities is a list of openapi.ActivityModel)
	for _, activity := range activities {
		newId = activity.GetId()

		newLatitude, err := strconv.ParseFloat(*activity.AddressApplied.Latitude, 64)
		//return if error raises while parsing
		if err != nil {
			return nil, err
		}
		newLongtitude, err := strconv.ParseFloat(*activity.AddressApplied.Longitude, 64)
		//return if error raises while parsing
		if err != nil {
			return nil, err
		}
		//create new coordinates using the long and lat we got from the activity
		newCoordinates := coordinates{
			latitude:   newLatitude,
			longtitude: newLongtitude,
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
