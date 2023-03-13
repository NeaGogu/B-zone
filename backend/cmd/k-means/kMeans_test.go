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

// Define some test data

func TestKMeans(t *testing.T) {
	// Define some test data
	activity1 := makeActivity(t, 1, "1", "1")
	activity2 := makeActivity(t, 1, "3", "4")
	activity3 := makeActivity(t, 1, "5", "6")
	activities := []openapi.ActivityModel{*activity1, *activity2, *activity3}

	testCases := []struct {
		name                string
		activities          []openapi.ActivityModel
		nrClusters          int
		nrCandidateClusters int
		expectedError       error
	}{
		{
			name:                "NrClusters smaller than 1",
			activities:          activities,
			nrClusters:          0,
			nrCandidateClusters: 2,
			expectedError:       ErrNrClustersTooSmall,
		},
		{
			name:                "NrCandidateClusters smaller or equal to 0",
			activities:          activities,
			nrClusters:          2,
			nrCandidateClusters: 0,
			expectedError:       ErrNrCandidateClustersTooSmall,
		},
		{
			name:                "Happy path",
			activities:          activities,
			nrClusters:          2,
			nrCandidateClusters: 2,
			expectedError:       nil,
		},
		{
			name:                "No activities",
			activities:          []openapi.ActivityModel{},
			nrClusters:          2,
			nrCandidateClusters: 2,
			expectedError:       ErrNoActivities,
		},
	}

	// Loop over test cases and run each one
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			got, err := kMeans(tc.activities, tc.nrClusters, tc.nrCandidateClusters)

			// Check the result
			if err != tc.expectedError {
				t.Errorf("kMeans() error = %v, want %v", err, tc.expectedError)
			}

			// If there was no error, check that the number of clusters matches the input
			if err == nil && len(got) != tc.nrClusters {
				t.Errorf("kMeans() returned %v clusters, want %v clusters", len(got), tc.nrClusters)
			}

			// Check that each cluster has at least one observation
			for _, c := range got {
				if len(c.observations) == 0 {
					t.Errorf("kMeans() returned a cluster with 0 observations")
				}
			}
		})
	}
}

func TestChooseCandidates(t *testing.T) {
	// Define some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3.0, 4.0), 2)
	obs3 := createObservation(t, createCoordinate(t, 5.0, 6.0), 3)
	observations := []observation{obs1, obs2, obs3}
	probabilities := []float64{0.2, 0.5, 0.3}
	nrCandidateClusters := 2

	// Seed the random number generator to ensure reproducibility
	randSeed := rand.New(rand.NewSource(12345))

	// Call the function and check the result
	want := []cluster{
		createCluster(t, obs3.coordinates, []observation{obs3}),
		createCluster(t, obs2.coordinates, []observation{obs2}),
	}
	got := chooseCandidates(observations, probabilities, nrCandidateClusters, randSeed)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n type of got1.observation %v, type of got2.observation %v\n type of want1.observation %v, type of want2.observation %v\n ", got, want, reflect.TypeOf(got[0].observations), reflect.TypeOf(got[1].observations), reflect.TypeOf(want[0].observations), reflect.TypeOf(want[1].observations))
	}
}

func TestProbabilities(t *testing.T) {
	// Define some test data
	distances := []float64{1.0, 2.0, 3.0}

	// Call the function and check the result
	want := []float64{1.0, 4.0, 9.0}
	got := Probabilities(distances)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Probabilities() = %v, want %v", got, want)
	}
}

func TestTotalSumDistance(t *testing.T) {
	// Define some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	obs3 := createObservation(t, createCoordinate(t, 5, 6), 1)
	observations := []observation{obs1, obs2, obs3}

	cluster1 := createCluster(t, obs1.coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.coordinates, []observation{obs2})
	clusters := []cluster{cluster1, cluster2}

	// Define a test table with different scenarios
	tests := []struct {
		name         string
		observations []observation
		checkCluster []cluster
		want         float64
		wantErr      bool
		err          error
	}{
		{
			name:         "normal case",
			observations: observations,
			checkCluster: clusters,
			want:         math.Sqrt(8),
			wantErr:      false,
			err:          nil,
		},
		{
			name:         "empty observations",
			observations: []observation{},
			checkCluster: clusters,
			want:         0,
			wantErr:      false,
			err:          nil,
		},
		{
			name:         "empty clusters",
			observations: observations,
			checkCluster: []cluster{},
			want:         0,
			wantErr:      true,
			err:          ErrNoClusters,
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := totalSumDistance(tt.observations, tt.checkCluster)

			if tt.wantErr && err == nil {
				t.Errorf("totalSumDistance() error = nil, wantErr = true")
			} else if !tt.wantErr && err != nil {
				t.Errorf("totalSumDistance() error = %v, wantErr = false", err)
			}

			if got != tt.want {
				t.Errorf("totalSumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalListDistance(t *testing.T) {
	// Define some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	obs3 := createObservation(t, createCoordinate(t, 5, 6), 1)
	observations := []observation{obs1, obs2, obs3}

	cluster1 := createCluster(t, obs1.coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.coordinates, []observation{obs2})
	clusters := []cluster{cluster1, cluster2}

	// Define a test table with different scenarios
	tests := []struct {
		name          string
		observations  []observation
		clusters      []cluster
		wantDistances []float64
		wantErr       bool
		err           error
	}{
		{
			name:          "normal case",
			observations:  observations,
			clusters:      clusters,
			wantDistances: []float64{math.Sqrt(8), 0, math.Sqrt(8)},
			wantErr:       false,
			err:           nil,
		},
		{
			//since the makes a list based on the length of observations and loops over observations we want to recieve an empty list back
			name:          "empty observations",
			observations:  []observation{},
			clusters:      clusters,
			wantDistances: []float64{},
			wantErr:       false,
			err:           nil,
		},
		{
			//because empty cluster defaults to center {0,0} we want the distances measured to 0
			name:          "empty clusters",
			observations:  observations,
			clusters:      []cluster{},
			wantDistances: []float64{2.23606797749979, 5, 7.810249675906654},
			wantErr:       false,
			err:           nil,
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotDistances []float64
			var err error
			if len(tt.clusters) > 0 {
				gotDistances, err = totalListDistance(tt.observations, tt.clusters[len(tt.clusters)-1])
			} else {
				gotDistances, err = totalListDistance(tt.observations, cluster{})
			}

			if tt.wantErr && err == nil {
				t.Errorf("totalListDistance() error = nil, wantErr = true")
			} else if !tt.wantErr && err != nil {
				t.Errorf("totalListDistance() error = %v, wantErr = false", err)
			}

			if !reflect.DeepEqual(gotDistances, tt.wantDistances) {
				t.Errorf("totalListDistance() distances = %v, want %v", gotDistances, tt.wantDistances)
			}
		})
	}
}

func TestBestCandidateFunc(t *testing.T) {
	// Create some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	observations := []observation{obs1, obs2}

	cluster1 := createCluster(t, obs1.coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.coordinates, []observation{obs2})
	clusters := []cluster{cluster1, cluster2}

	candidate1 := cluster{center: coordinates{latitude: 5.0, longitude: 6.0}, observations: []observation{}}
	candidate2 := cluster{center: coordinates{latitude: 7.0, longitude: 8.0}, observations: []observation{}}
	candidates := []cluster{candidate1, candidate2}

	t.Run("Normal run", func(t *testing.T) {
		// Call the function and check the result
		want := candidate1
		got, err := bestCandidateFunc(candidates, clusters, observations)
		assertEqual(t, got, want, err)
	})

	t.Run("Two candidates have the same distance", func(t *testing.T) {
		// Create some test data
		obsEqual1 := createObservation(t, createCoordinate(t, 1, 2), 1)
		obsEqual2 := createObservation(t, createCoordinate(t, 2, 1), 2)
		observationsEqual := []observation{obsEqual1, obsEqual2}
		//create two candidates that have equal sum of distances
		middleCandidate1 := createCluster(t, createCoordinate(t, 1, 2), []observation{})
		middleCandidate2 := createCluster(t, createCoordinate(t, 2, 1), []observation{})
		candidates := []cluster{middleCandidate1, middleCandidate2}
		//since the comparison is < we expect the first candidate if the result is equal
		want := middleCandidate1
		// Call the function and check the result
		got, err := bestCandidateFunc(candidates, []cluster{}, observationsEqual)
		assertEqual(t, got, want, err)
	})
	t.Run("No candidates given", func(t *testing.T) {
		candidates := []cluster{}
		_, err := bestCandidateFunc(candidates, clusters, observations)
		assertError(t, err, ErrNrCandidateClustersTooSmall)
	})

}
func TestDistanceToNearestCluster(t *testing.T) {
	//Create clusters to test with
	clusters := clusters{
		createCluster(t, createCoordinate(t, 4, 6), make(observations, 0)),
		createCluster(t, createCoordinate(t, 0, 0), make(observations, 0)),
	}
	// GIS/handheld gps precision
	epsilon := 0.00001
	t.Run("Center difference on longitude", func(t *testing.T) {
		// Create observation with difference on the longitude and compare if distance and cluster index are correct
		observation := createObservation(t, createCoordinate(t, 2, 6), 1)
		gotDistance, gotIndex, err := distanceToNearestCluster(observation, clusters)
		wantDistance := 2.0
		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}
		result := AlmostEqual(t, gotDistance, wantDistance, epsilon)
		if !result {
			t.Errorf("The difference between gotDistance: %f and wantDistance: %f is too large", gotDistance, wantDistance)
		}
		wantIndex := 0
		if !(gotIndex == wantIndex) {
			t.Errorf("The selected cluster :%v is not the closest cluster %v", gotIndex, wantIndex)
		}
	})
	t.Run("Center difference on latitude", func(t *testing.T) {
		// Create observation with difference on the latitude and compare if distance and cluster index are correct
		observation := createObservation(t, createCoordinate(t, 4, 8), 1)
		gotDistance, gotIndex, err := distanceToNearestCluster(observation, clusters)
		wantDistance := 2.0
		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}
		result := AlmostEqual(t, gotDistance, wantDistance, epsilon)
		if !result {
			t.Errorf("The difference between gotDistance: %f and wantDistance: %f is too large", gotDistance, wantDistance)
		}
		wantIndex := 0
		if !(gotIndex == wantIndex) {
			t.Errorf("The selected cluster :%v is not the closest cluster %v", gotIndex, wantIndex)
		}
	})
	t.Run("point is in the middle of two centers", func(t *testing.T) {
		// Create observation that is exaclty in the middle compare if distance and cluster index are correct
		observation := createObservation(t, createCoordinate(t, 2, 3), 1)
		gotDistance, gotIndex, err := distanceToNearestCluster(observation, clusters)
		wantDistance := 3.605551275464
		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}
		result := AlmostEqual(t, gotDistance, wantDistance, epsilon)
		if !result {
			t.Errorf("The difference between gotDistance: %f and wantDistance: %f is too large", gotDistance, wantDistance)
		}
		wantIndex := 0
		if !(gotIndex == wantIndex) {
			t.Errorf("The selected cluster :%v is not the closest cluster %v", gotIndex, wantIndex)
		}
	})
	t.Run("point is on center", func(t *testing.T) {
		// Create observation that is exaclty on the cluster and compare if distance and cluster index are correct
		observation := createObservation(t, createCoordinate(t, 0, 0), 1)
		gotDistance, gotIndex, err := distanceToNearestCluster(observation, clusters)
		wantDistance := 0.0
		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}
		result := AlmostEqual(t, gotDistance, wantDistance, epsilon)
		if !result {
			t.Errorf("The difference between gotDistance: %f and wantDistance: %f is too large", gotDistance, wantDistance)
		}
		wantIndex := 1
		if !(gotIndex == wantIndex) {
			t.Errorf("The selected cluster :%v is not the closest cluster %v", gotIndex, wantIndex)
		}
	})
}

func TestDistance(t *testing.T) {

	t.Run("easy run", func(t *testing.T) {
		//create an observation and a cluster with known coordinates
		observation := createObservation(t, createCoordinate(t, 10, 20), 1)
		cluster := createCluster(t, createCoordinate(t, 15, 25), make(observations, 0))
		//calculate the expected distance between the observation and the cluster
		expectedDistance := math.Sqrt((observation.coordinates.latitude-cluster.center.latitude)*(observation.coordinates.latitude-cluster.center.latitude) +
			(observation.coordinates.longitude-cluster.center.longitude)*(observation.coordinates.longitude-cluster.center.longitude))

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})
	t.Run("latitude negative", func(t *testing.T) {
		//create an observation and a cluster with known coordinates
		observation := createObservation(t, createCoordinate(t, -5, 5), 1)
		cluster := createCluster(t, createCoordinate(t, 5, 5), make(observations, 0))
		//calculate the expected distance between the observation and the cluster
		expectedDistance := 10.0

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})
	t.Run("longitude negative", func(t *testing.T) {
		//create an observation and a cluster with known coordinates
		observation := createObservation(t, createCoordinate(t, 5, -5), 1)
		cluster := createCluster(t, createCoordinate(t, 5, 5), make(observations, 0))
		//calculate the expected distance between the observation and the cluster
		expectedDistance := 10.0

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})
	t.Run("complex float", func(t *testing.T) {
		//create an observation and a cluster with known coordinates
		observation := createObservation(t, createCoordinate(t, 51.4486098, 5.4907148), 1)
		cluster := createCluster(t, createCoordinate(t, 51.395216, 5.474121), make(observations, 0))
		//calculate the expected distance between the observation and the cluster
		expectedDistance := math.Sqrt((observation.coordinates.latitude-cluster.center.latitude)*(observation.coordinates.latitude-cluster.center.latitude) +
			(observation.coordinates.longitude-cluster.center.longitude)*(observation.coordinates.longitude-cluster.center.longitude))

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0000001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})

	t.Run("distance with 0", func(t *testing.T) {
		//create an observation and a cluster with known coordinates
		observation := createObservation(t, createCoordinate(t, 2, 3), 1)
		cluster := createCluster(t, createCoordinate(t, 0, 0), make(observations, 0))
		//calculate the expected distance between the observation and the cluster
		expectedDistance := 3.605551275464

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0000001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})

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
				latitude:  51.5074,
				longitude: -0.1278,
			},
		},
		{
			id: 2,
			coordinates: coordinates{
				latitude:  52.5200,
				longitude: 13.4050,
			},
		},
	}

	var observations = make(observations, 0)
	got, err := activitiesToObservations(activities, observations)

	assertEqual(t, got, want, err)

}

func TestInitializeClusters(t *testing.T) {
	// Define some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	obs3 := createObservation(t, createCoordinate(t, 5, 6), 3)
	observations := []observation{obs1, obs2, obs3}

	testCases := []struct {
		name                string
		nrClusters          int
		nrCandidateClusters int
		wantError           bool
		expectedError       error
		expectedClusters    clusters
	}{
		{
			name:                "NrClusters larger than length of observations",
			nrClusters:          4,
			nrCandidateClusters: 2,
			expectedError:       ErrNrObservationsTooSmall,
			wantError:           true,
			expectedClusters:    clusters{},
		},
		{
			name:                "NrCandidateClusters smaller or equal to 0",
			nrClusters:          2,
			nrCandidateClusters: 0,
			expectedError:       ErrNrCandidateClustersTooSmall,
			wantError:           true,
			expectedClusters:    clusters{},
		},
		{
			name:                "NrCandidateClusters larger than length of observations",
			nrClusters:          2,
			nrCandidateClusters: 4,
			expectedError:       ErrNrObservationsTooSmall,
			wantError:           true,
			expectedClusters:    clusters{},
		},
		{
			name:                "NrClusters equal to length of observations",
			nrClusters:          3,
			nrCandidateClusters: 2,
			expectedError:       nil,
			wantError:           false,
			expectedClusters: clusters{
				createCluster(t, obs3.coordinates, []observation{obs3}),
				createCluster(t, obs1.coordinates, []observation{obs1}),
				createCluster(t, obs2.coordinates, []observation{obs2}),
			},
		},
		{
			name:                "Happy path",
			nrClusters:          2,
			nrCandidateClusters: 2,
			expectedError:       nil,
			wantError:           false,
			expectedClusters: clusters{
				createCluster(t, obs2.coordinates, []observation{obs2}),
				createCluster(t, obs3.coordinates, []observation{obs3}),
			},
		},
	}

	// Loop over test cases and run each one
	// Seed the random number generator to ensure reproducibility
	randSeed := rand.New(rand.NewSource(12345))
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			clusters := clusters{}
			got, err := initializeClusters(observations, clusters, tc.nrClusters, tc.nrCandidateClusters, randSeed)

			// Check the result
			if err != tc.expectedError {
				t.Errorf("initializeClusters() error = %v, want %v", err, tc.expectedError)
			}
			if !tc.wantError && !reflect.DeepEqual(got, tc.expectedClusters) {
				t.Errorf("initializeClusters() = %v, want %v", got, tc.expectedClusters)
			}
		})
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

func createCluster(t testing.TB, center coordinates, observations observations) cluster {
	t.Helper()
	cluster := cluster{
		center:       center,
		observations: observations,
	}
	return cluster
}

func createObservation(t testing.TB, coordinates coordinates, id int64) observation {
	t.Helper()
	observation := observation{
		id:          id,
		coordinates: coordinates,
	}
	return observation
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
		latitude:  lat,
		longitude: long,
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
