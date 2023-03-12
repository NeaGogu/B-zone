package kMeans

import (
	openapi "backend/internal/swag_gen"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

// MinNormal is the smallest positive normal value of type float64.
var MinNormal = math.Float64frombits(0x0010000000000000)

func TestKMeans(t *testing.T) {
	// Initialize test data
	activity1 := makeActivity(t, 1, "1.234567", "2.345678")

	activity2 := makeActivity(t, 2, "3.456789", "4.567890")
	activity3 := makeActivity(t, 3, "5.678901", "6.789012")

	activities := activities{
		*activity1,
		*activity2,
		*activity3,
	}

	t.Run("Test kMeans function with 2 clusters", func(t *testing.T) {
		clusters, err := kMeans(activities, 2, 2)

		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}

		if len(clusters) != 2 {
			t.Errorf("Expected 2 clusters, got %d", len(clusters))
		}
	})

	t.Run("Test kMeans function with 3 clusters", func(t *testing.T) {
		clusters, err := kMeans(activities, 3, 2)
		//t.Errorf("clusters: %v", clusters)

		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}

		t.Errorf("%v", clusters)
		if len(clusters) != 3 {
			t.Errorf("Expected 3 clusters, got %d", len(clusters))
		}
	})

}

func TestBestCandidateFunc(t *testing.T) {
	// Create some test data
	obs1 := observation{id: 1, coordinates: coordinates{latitude: 1.0, longtitude: 2.0}}
	obs2 := observation{id: 2, coordinates: coordinates{latitude: 3.0, longtitude: 4.0}}
	observations := []observation{obs1, obs2}

	cluster1 := cluster{center: obs1.coordinates, observations: []observation{obs1}}
	cluster2 := cluster{center: obs2.coordinates, observations: []observation{obs2}}
	clusters := []cluster{cluster1, cluster2}

	candidate1 := cluster{center: coordinates{latitude: 5.0, longtitude: 6.0}, observations: []observation{}}
	candidate2 := cluster{center: coordinates{latitude: 7.0, longtitude: 8.0}, observations: []observation{}}
	candidates := []cluster{candidate1, candidate2}

	t.Run("Normal run", func(t *testing.T) {
		// Call the function and check the result
		want := candidate1
		got := bestCandidateFunc(candidates, clusters, observations)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("bestCandidateFunc() = %v, want %v", got, want)
		}
	})

	t.Run("Run in the middle", func(t *testing.T) {
		middleCandidate := cluster{center: coordinates{latitude: 2.0, longtitude: 3.0}, observations: []observation{}}
		middleCandidateCopy := cluster{center: coordinates{latitude: 2.0, longtitude: 3.0}, observations: []observation{}}
		candidates := []cluster{middleCandidate, middleCandidateCopy}
		//since the comparison is < we expect the first candidate if the result is equal
		want := middleCandidate
		// Call the function and check the result
		got := bestCandidateFunc(candidates, clusters, observations)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("bestCandidateFunc() = %v, want %v", got, want)
		}
	})

}
func TestDistanceToNearestCluster(t *testing.T) {
	clusters := clusters{
		cluster{center: coordinates{longtitude: 4, latitude: 6}, observations: make(observations, 0)},
		cluster{center: coordinates{longtitude: 0, latitude: 0}, observations: make(observations, 0)},
	}
	//GIS/ handheld gps precision
	epsilon := 0.00001
	t.Run("Center difference on longtitude", func(t *testing.T) {
		observation := observation{id: 1, coordinates: coordinates{longtitude: 2, latitude: 6}}
		got, _ := distanceToNearestCluster(observation, clusters)
		want := 2.0
		result := AlmostEqual(t, got, want, epsilon)
		if !result {
			t.Errorf("The difference between got: %f and want: %f is too large", got, want)
		}
	})
	t.Run("Center difference on latitude", func(t *testing.T) {
		observation := observation{id: 1, coordinates: coordinates{longtitude: 4, latitude: 8}}
		got, _ := distanceToNearestCluster(observation, clusters)
		want := 2.0
		result := AlmostEqual(t, got, want, epsilon)
		if !result {
			t.Errorf("The difference between got: %f and want: %f is too large", got, want)
		}
	})
	t.Run("point is in the middle of two centers", func(t *testing.T) {
		observation := observation{id: 1, coordinates: coordinates{longtitude: 2, latitude: 3}}
		got, _ := distanceToNearestCluster(observation, clusters)
		want := 3.605551275464
		result := AlmostEqual(t, got, want, epsilon)
		if !result {
			t.Errorf("The difference between got: %f and want: %f is too large", got, want)
		}
	})
	t.Run("point is on center", func(t *testing.T) {
		observation := observation{id: 1, coordinates: coordinates{longtitude: 0, latitude: 0}}
		got, _ := distanceToNearestCluster(observation, clusters)
		want := 0.0
		result := AlmostEqual(t, got, want, epsilon)
		if !result {
			t.Errorf("The difference between got: %f and want: %f is too large", got, want)
		}
	})
}

func TestDistance(t *testing.T) {
	//create an observation and a cluster with known coordinates
	observation := observation{
		id: 1,
		coordinates: coordinates{
			latitude:   10,
			longtitude: 20,
		},
	}
	cluster := cluster{
		center: coordinates{
			latitude:   15,
			longtitude: 25,
		},
	}

	//calculate the expected distance between the observation and the cluster
	expectedDistance := math.Sqrt((observation.coordinates.latitude-cluster.center.latitude)*(observation.coordinates.latitude-cluster.center.latitude) +
		(observation.coordinates.longtitude-cluster.center.longtitude)*(observation.coordinates.longtitude-cluster.center.longtitude))

	//call the distance function to calculate the actual distance
	actualDistance := distance(observation, cluster)

	//check if the actual distance matches the expected distance
	if actualDistance != expectedDistance {
		t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
	}
}

func TestActivitiesToObservations(t *testing.T) {
	activity1 := makeActivity(t, 1, "51.5074", "-0.1278")
	activity2 := makeActivity(t, 2, "52.5200", "13.4050")
	activities := activities{
		*activity1,
		*activity2,
	}
	want := observations{
		{
			id: 1,
			coordinates: coordinates{
				latitude:   51.5074,
				longtitude: -0.1278,
			},
		},
		{
			id: 2,
			coordinates: coordinates{
				latitude:   52.5200,
				longtitude: 13.4050,
			},
		},
	}

	var observations = make(observations, 0)
	got, err := activitiesToObservations(activities, observations)

	assertEqual(t, got, want, err)

}

func TestInitializeClusters(t *testing.T) {
	// Create some dummy observations
	observations := make(observations, 0)
	for i := 0; i < 10; i++ {
		lat := rand.Float64() * 10
		long := rand.Float64() * 10
		observations = append(observations, observation{
			id: int64(i),
			coordinates: coordinates{
				latitude:   lat,
				longtitude: long,
			},
		})
	}

	// Initialize clusters with 3 clusters
	nrClusters := 3
	nrCandidateClusters := 5
	clusters := make(clusters, 0, nrClusters)
	clusters, err := initializeClusters(observations, clusters, nrClusters, nrCandidateClusters)

	// Check that there is no error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check that the number of clusters is correct
	if len(clusters) != nrClusters {
		t.Errorf("Expected %d clusters, but got %d", nrClusters, len(clusters))
	}

	// Check that each cluster has at least one observation
	for _, cluster := range clusters {
		if len(cluster.observations) < 1 {
			t.Errorf("Expected each cluster to have at least one observation, but one cluster has %d observations", len(cluster.observations))
		}
	}

	// Check that each observation is assigned to a cluster
	for _, observation := range observations {
		assigned := false
		for _, cluster := range clusters {
			for _, obs := range cluster.observations {
				if observation.id == obs.id {
					assigned = true
					break
				}
			}
			if assigned {
				break
			}
		}
		if !assigned {
			t.Errorf("Expected observation with id %d to be assigned to a cluster, but it wasn't", observation.id)
		}
	}
}

func BenchmarkKMeansALot(t *testing.B) {
	activities := make(activities, 0)
	for i := 0; i < 100000; i++ {
		str1 := strconv.Itoa(rand.Intn(1000))
		str2 := strconv.Itoa(rand.Intn(1000))
		result := makeActivity(t, int64(i), str1, str2)
		activities = append(activities, *result)
	}
	kMeans(activities, 30, 10)
}

// AlmostEqual returns true if a and b are equal within a relative error of
// ε. See http://floating-point-gui.de/errors/comparison/ for the details of the
// applied method.
func AlmostEqual(t testing.TB, a, b, ε float64) bool {
	t.Helper()
	if a == b {
		return true
	}
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)
	if a == 0 || b == 0 || absA+absB < MinNormal {
		return diff < ε*MinNormal
	}
	return diff/math.Min(absA+absB, math.MaxFloat64) < ε
}

func createCoordinate(t testing.TB, lat float64, long float64) coordinates {
	t.Helper()
	coord := coordinates{
		latitude:   lat,
		longtitude: long,
	}

	return coord
}

func createObservations(t testing.TB, coordinatesList []coordinates) observations {
	t.Helper()

	observations := make(observations, 0)

	for i, coord := range coordinatesList {
		observations = append(observations, observation{
			id:          int64(i),
			coordinates: coord,
		})
	}

	return observations
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Errorf("didn't get an error but wanted %q", want)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertEqual[T any](t testing.TB, got, want T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("didn't want an error but got: %q", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v \n want %v", got, want)
	}
}

func makeActivity(t testing.TB, id int64, lat string, long string) *openapi.ActivityModel {
	t.Helper()

	//make activity and addressapplied model
	activity := openapi.NewActivityModel()
	address := openapi.NewAddressAppliedModel()
	//set lat and long
	address.SetLatitude(lat)
	address.SetLongitude(long)
	activity.SetAddressApplied(*address)
	activity.SetId(id)

	return activity
}
