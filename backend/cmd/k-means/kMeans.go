package kMeans

import (
	"errors"
	"math"
	"math/rand"
	"reflect"
	"strconv"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/mroth/weightedrand/v2"
)

type coordinates struct {
	long float64
	lat  float64
}

type centroid struct {
	center             coordinates
	AssignedActivities []openapi.ActivityModel
}

var ErrNrActivitiesTooSmall = errors.New("number of activities smaller than number of clusters")
var ErrNrCandidateCentersTooSmall = errors.New("number of candidate smaller than or equal to 0")
var ErrNoActivities = errors.New("number of candidate smaller than or equal to 0")

// refactor since it is used multiple times
var emptyCentroid = []centroid{
	{
		center: coordinates{
			long: 0,
			lat:  0,
		},
		AssignedActivities: []openapi.ActivityModel{},
	},
}

func kMeans(activities []openapi.ActivityModel, nrClusters int) ([]centroid, error) {

	//initialize centroids
	centroids, err := initCentroids(activities, nrClusters, 2)
	if err != nil {
		return centroids, err
	}
	converged := false

	//stop if converged
	for !converged {
		oldCentroids := centroids

		//clear all the assigned activities
		centroids, err = clearActivities(centroids)

		if err != nil {
			return centroids, err
		}

		for _, activity := range activities {
			//for each activity calculate the nearest centroid and assign activity to that centroid
			closestCluster := nearestCentroid(activity, centroids)
			centroids[closestCluster].AssignedActivities = append(centroids[closestCluster].AssignedActivities, activity)
		}
		//update the clusters

		centroids, err := updateCluster(centroids)
		if err != nil {
			return centroids, err
		}

		if reflect.DeepEqual(oldCentroids, centroids) {
			converged = true
		}

	}

	//listAssignedActivities := listActivities()
	return centroids, nil
}

// initializes cluster by greedy k-means++ seeding
func initCentroids(activities []openapi.ActivityModel, nrClusters int, nrCandidateCenters int) ([]centroid, error) {

	//throw error if nr of activities is smaller than or equal to 0
	if len(activities) <= 0 {
		return emptyCentroid, ErrNoActivities
	}
	//throw error if nrCluster is larger than length of activities or nrCandidateCenters smaller or equal to 0
	if nrClusters > len(activities) {
		return emptyCentroid, ErrNrActivitiesTooSmall
	}
	//throw error if nr of candidate centers is smaller or equal to 0
	if nrCandidateCenters <= 0 {
		return emptyCentroid, ErrNrCandidateCentersTooSmall
	}

	//initialize centroids
	centroids := make([]centroid, 0, nrClusters)
	//intialize slice to hold candidate centers
	var candidateCenters []centroid
	var err error
	//Sample l candidateCenters
	candidateCenters = make([]centroid, 0, nrCandidateCenters)
	for i := 0; i < nrCandidateCenters; i++ {
		randomInt := rand.Intn(len(activities))

		//sample candidates using random number (only done for first centroid)
		candidateCenters, err = sampleCandidateCentroid(activities, randomInt, candidateCenters)

		if err != nil {
			return candidateCenters, err
		}
	}

	//calculate iteration with lowest cost
	lowestCost := lowestcandidateCenterCost(activities, nrCandidateCenters, candidateCenters)

	// long, _ := strconv.ParseFloat(*activities[lowestCost].AddressApplied.Longitude, 64)
	// lat, _ := strconv.ParseFloat(*activities[lowestCost].AddressApplied.Latitude, 64)

	//append lowestcost centroid
	addedCentroid := candidateCenters[lowestCost]

	centroids = append(centroids, addedCentroid)

	//Iteratively choose best cluster to add
	for i := 0; i < nrClusters-1; i++ {
		//clear old candidateCenters
		candidateCenters = make([]centroid, 0, nrCandidateCenters)
		//calculate bias of points with higher bias towards points further away from previous centroid
		choices := newChoices(activities, centroids[i])
		//make random chooser with bias
		chs, err := weightedrand.NewChooser(choices...)

		if err != nil {
			return emptyCentroid, err
		}

		//choose new candidate center with bias towards points furtheraway
		for i := 0; i < nrCandidateCenters; i++ {
			biasRandomInt := chs.Pick()
			candidateCenters, err = sampleCandidateCentroid(activities, biasRandomInt, candidateCenters)

			if err != nil {
				return candidateCenters, err
			}
		}
		//calculate candidateCenter with lowest cost
		lowestCost = lowestcandidateCenterCost(activities, nrCandidateCenters, candidateCenters)

		//append lowestcost centroid
		centroids = append(centroids, candidateCenters[lowestCost])
	}
	return centroids, nil
}

// Samples a candidates from activities given a random (or biased) int
func sampleCandidateCentroid(activities []openapi.ActivityModel, randomInt int, candidateCenter []centroid) ([]centroid, error) {
	long, err := strconv.ParseFloat(*activities[randomInt].AddressApplied.Longitude, 64)
	//return if converting throws error
	if err != nil {
		return emptyCentroid, err
	}

	lat, err := strconv.ParseFloat(*activities[randomInt].AddressApplied.Latitude, 64)

	//return if converting throws error
	if err != nil {
		return emptyCentroid, err
	}

	candidateCenter = append(candidateCenter, centroid{
		center: coordinates{
			long: long,
			lat:  lat,
		},
		AssignedActivities: []openapi.ActivityModel{activities[randomInt]},
	})

	return candidateCenter, nil
}

// function to calculate which candidate center has the lowest cost (cummilative distance)
func lowestcandidateCenterCost(activities []openapi.ActivityModel, nrCandidateCenters int, candidateCenters []centroid) int {
	var lowIterationCost = 0
	minCost := math.Inf(0)
	for i := 0; i < nrCandidateCenters; i++ {
		var iterationCost float64

		for j := 0; j < len(activities); j++ {
			cost := distanceToCluster(candidateCenters[i], activities[i])
			iterationCost += cost
		}
		if iterationCost < minCost {
			minCost = iterationCost
			lowIterationCost = i
		}
	}
	return lowIterationCost
}

func newChoices(activities []openapi.ActivityModel, centroid centroid) []weightedrand.Choice[int, int] {
	//slice to store distance to given centroid used for sampling with probability
	distanceToCentroid := make([]float64, len(activities))
	distanceSum := 0.0

	//choices
	choices := make([]weightedrand.Choice[int, int], len(activities))
	//calculate distance to centroid
	for index := range activities {
		distanceToCentroid[index] = distanceToCluster(centroid, activities[index])
		distanceSum += distanceToCentroid[index]
	}

	//multiply coordinates by 1000 to get a precision of 10 meters
	for index := range activities {
		distanceToCentroid[index] = distanceToCentroid[index] / distanceSum
		tempChoice := weightedrand.NewChoice(index, int(math.Round(distanceToCentroid[index]*1000)))
		choices[index] = tempChoice
	}
	return choices
}

// Clears all activities from given centroids
func clearActivities(centroids []centroid) ([]centroid, error) {
	for i := range centroids {
		centroids[i].AssignedActivities = make([]openapi.ActivityModel, 0)
	}
	return centroids, nil
}

// Calculates and returns the nearest cluster
func nearestCentroid(activity openapi.ActivityModel, clusters []centroid) int {
	lowestDistance := math.Inf(0)
	var indexCluster int
	for i, j := range clusters {
		distance := distanceToCluster(j, activity)
		if distance < lowestDistance {
			lowestDistance = distance
			indexCluster = i
		}
	}
	return indexCluster
}

// Calculates the center of the cluster (by average of all assigned activities)
func updateCluster(clusters []centroid) ([]centroid, error) {
	for i := range clusters {
		var long float64 = 0
		var lat float64 = 0
		for _, j := range clusters[i].AssignedActivities {
			tempLong, err := strconv.ParseFloat(*j.AddressApplied.Longitude, 64)
			if err != nil {
				return clusters, err
			}
			tempLat, err := strconv.ParseFloat(*j.AddressApplied.Latitude, 64)
			if err != nil {
				return clusters, err
			}
			long += tempLong
			lat += tempLat
		}

		//update center
		clusters[i].center.lat = lat / float64(len(clusters[i].AssignedActivities))
		clusters[i].center.long = long / float64(len(clusters[i].AssignedActivities))
	}
	return clusters, nil
}

// puts all clusters in a list of list so it can be made into a zone
func listActivities() []int {
	return []int{0}
}

// Returns the distance to given cluster
func distanceToCluster(centroid centroid, activity openapi.ActivityModel) float64 {
	longCluster := centroid.center.long
	latCluster := centroid.center.lat
	longActivity, _ := strconv.Atoi(*activity.AddressApplied.Longitude)
	latActivity, _ := strconv.Atoi(*activity.AddressApplied.Latitude)

	distance := math.Sqrt(math.Pow(longCluster-float64(longActivity), 2) + math.Pow(latCluster-float64(latActivity), 2))
	return distance

}
