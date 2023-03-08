package kMeans

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/mroth/weightedrand/v2"
)

type coordinate struct {
	coordinate float64
}

type coordinates struct {
	long float64
	lat  float64
}

type centroid struct {
	center             coordinates
	AssignedActivities openapi.ActivityModel
}

var ErrNrActivitiesTooSmall = errors.New("Number of activities smaller than number of clusters")
var ErrNrCandidateCentersTooSmall = errors.New("Number of candidate smaller than or equal to 0")
var ErrNoActivities = errors.New("Number of candidate smaller than or equal to 0")

func kMeans(activities []openapi.ActivityModel, nrClusters int) ([][]openapi.ActivityModel, error) {

	//initialize centroids
	_, err := initCentroids(activities, nrClusters, 10)
	if err != nil {
		return [][]openapi.ActivityModel{
			{{}, {}},
			{{}, {}},
		}, err
	}
	converged := false

	//stop if converged
	for !converged {
		for range activities {
			//clear all the assigned activities
			clearActivities()
			//for each activity calculate the nearest centroid and assign activity to that centroid
			//closestCluster := nearestCluster()
			//centroids[closestCluster].AssignedActivities += activities
		}
		//update the clusters
		updateCluster()
		converged = true
	}

	//listAssignedActivities := listActivities()
	return [][]openapi.ActivityModel{
		{}, {}}, fmt.Errorf("test")
}

// initializes cluster by greedy k-means++ seeding
func initCentroids(activities []openapi.ActivityModel, nrClusters int, nrCandidateCenters int) ([]centroid, error) {

	//throw error if nr of activities is smaller than or equal to 0
	if len(activities) <= 0 {
		return []centroid{
			{
				center: coordinates{
					long: 0,
					lat:  0,
				},
				AssignedActivities: openapi.ActivityModel{},
			},
		}, ErrNoActivities
	}
	//throw error if nrCluster is larger than length of activities or nrCandidateCenters smaller or equal to 0
	if nrClusters > len(activities) {
		return []centroid{
			{
				center: coordinates{
					long: 0,
					lat:  0,
				},
				AssignedActivities: activities[0],
			},
		}, ErrNrActivitiesTooSmall
	}
	//throw error if nr of candidate centers is smaller or equal to 0
	if nrCandidateCenters <= 0 {
		return []centroid{
			{
				center: coordinates{
					long: 0,
					lat:  0,
				},
				AssignedActivities: activities[0],
			},
		}, ErrNrCandidateCentersTooSmall
	}

	//initialize centroids
	centroids := make([]centroid, 1, nrClusters)
	//intialize slice to hold candidate centers
	candidateCenters := make([]centroid, 0, nrCandidateCenters)

	//Sample l candidateCenters
	for i := 0; i < nrCandidateCenters; i++ {
		randomInt := rand.Intn(len(activities))
		long, err := strconv.ParseFloat(*activities[randomInt].AddressApplied.Longitude, 64)
		lat, err := strconv.ParseFloat(*activities[randomInt].AddressApplied.Latitude, 64)

		//return if converting throws error
		if err != nil {
			return []centroid{
				centroid{
					center: coordinates{
						long: 0,
						lat:  0,
					},
					AssignedActivities: activities[0],
				},
			}, err
		}

		candidateCenters = append(candidateCenters, centroid{
			center: coordinates{
				long: long,
				lat:  lat,
			},
			AssignedActivities: activities[i],
		})
	}

	//variable to store the iteration with lowest cost
	var lowestIterationCost int

	//calculate iteration with lowest cost
	for i := 0; i < nrCandidateCenters; i++ {

		minCost := math.Inf(0)
		var iterationCost float64
		for j := 0; j < len(activities); j++ {
			cost := distanceToCluster(candidateCenters[i], activities[i])
			iterationCost += cost
		}
		if iterationCost < minCost {
			minCost = iterationCost
			lowestIterationCost = i
		}
	}

	centroids[0] = candidateCenters[lowestIterationCost]

	//Iteratively choose best cluster to add
	for i := 0; i < nrClusters-1; i++ {
		choices := newChoices(activities, centroids[i])
		chs, _ := weightedrand.NewChooser(choices...)
		for i := 0; i < nrCandidateCenters; i++ {
			biasRandomInt := chs.Pick()
			long, err := strconv.ParseFloat(*activities[biasRandomInt].AddressApplied.Longitude, 64)
			lat, err := strconv.ParseFloat(*activities[biasRandomInt].AddressApplied.Latitude, 64)

			//return if converting throws error
			if err != nil {
				return centroids, err
			}

			candidateCenters[i] = centroid{
				center: coordinates{
					long: long,
					lat:  lat,
				},
				AssignedActivities: activities[i],
			}
		}

		sampleCandidateCentroids(activities, 2, func(chs *weightedrand.Chooser[int, int]) int {
			return chs.Pick()
		})

		//calculate iteration with lowest cost
		for i := 0; i < nrCandidateCenters; i++ {

			minCost := math.Inf(0)
			var iterationCost float64
			for j := 0; j < len(activities); j++ {
				cost := distanceToCluster(candidateCenters[i], activities[i])
				iterationCost += cost
			}
			if iterationCost < minCost {
				minCost = iterationCost
				lowestIterationCost = i
			}
		}
		centroids = append(centroids, candidateCenters[lowestIterationCost])
	}
	return centroids, fmt.Errorf("weet nog niet welke error")
}

func newChoices(activities []openapi.ActivityModel, centroid centroid) []weightedrand.Choice[int, int] {
	//slice to store distance to given centroid used for sampling with probability
	distanceToCentroid := make([]float64, 0, len(activities))
	distanceSum := 0.0

	//choices
	choices := make([]weightedrand.Choice[int, int], 0, len(activities))
	//calculate distance to centroid
	for index, _ := range activities {
		distanceToCentroid[index] = distanceToCluster(centroid, activities[index])
		distanceSum += distanceToCentroid[index]
	}

	for index, _ := range activities {
		distanceToCentroid[index] = distanceToCentroid[index] / distanceSum
	}

	for index, _ := range activities {
		tempChoice := weightedrand.NewChoice(index, int(math.Round(distanceToCentroid[index])))
		choices = append(choices, tempChoice)
	}

	return choices
}

func sampleCandidateCentroids[T any](activities []openapi.ActivityModel, nrCandidateCenters int, random T) {

}

// Clears all activities from given centroids
func clearActivities() {

}

// Calculates and returns the nearest cluster
func nearestCluster() int {
	return 1
}

// Calculates the center of the cluster (by average of all assigned activities)
func updateCluster() {

}

// puts all clusters in a list of list so it can be made into a zone
func listActivities() []int {
	return []int{0}
}

// Returns the distance to given cluster
func distanceToCluster(centroid centroid, activity openapi.ActivityModel) float64 {
	return 1.00
}
