package kMeans

import (
	model "bzone/backend/internal/models"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MinNormal is the smallest positive normal value of type float64.
var MinNormal = math.Float64frombits(0x0010000000000000)

func TestKMeans(t *testing.T) {
	// Define some test data
	activity1 := makeActivity(t, "1", "1", "1", "1234")
	activity2 := makeActivity(t, "1", "3", "4", "1234")
	activity3 := makeActivity(t, "1", "5", "6", "1234")
	activities := []model.ActivityModelBumbal{*activity1, *activity2, *activity3}

	// Define test cases
	testTable := []struct {
		name                string
		activities          []model.ActivityModelBumbal
		nrClusters          int
		nrCandidateClusters int
		wantError           bool
		expectedError       error
	}{
		{
			name:                "NrClusters smaller than 1",
			activities:          activities,
			nrClusters:          0,
			nrCandidateClusters: 2,
			wantError:           true,
			expectedError:       ErrNrClustersTooSmall,
		},
		{
			name:                "NrCandidateClusters smaller or equal to 0",
			activities:          activities,
			nrClusters:          2,
			nrCandidateClusters: 0,
			wantError:           true,
			expectedError:       ErrNrCandidateClustersTooSmall,
		},
		{
			name:                "Happy path",
			activities:          activities,
			nrClusters:          2,
			nrCandidateClusters: 2,
			wantError:           false,
			expectedError:       nil,
		},
		{
			name:                "No activities",
			activities:          []model.ActivityModelBumbal{},
			nrClusters:          2,
			nrCandidateClusters: 2,
			wantError:           true,
			expectedError:       ErrNoActivities,
		},
	}

	// Loop over test cases and run each one
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function
			got, err := KMeans(tt.activities, tt.nrClusters, tt.nrCandidateClusters)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// If there was no error, check that the number of Clusters matches the input
			if !tt.wantError && len(got) != tt.nrClusters {
				t.Errorf("kMeans() returned %v Clusters, want %v Clusters", len(got), tt.nrClusters)
			}
		})
	}
}

func TestUpdateCluster(t *testing.T) {
	// Define test cases
	testTable := []struct {
		name     string
		Clusters Clusters
		want     Clusters
	}{
		{
			name: "Happy path",
			Clusters: Clusters{
				createCluster(t, createCoordinate(t, 52.0, 4.0), observations{
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 1),
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 2),
				}),
				createCluster(t, createCoordinate(t, 51.0, 4.3), observations{
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 3),
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 4),
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 5),
				}),
			},
			want: Clusters{
				createCluster(t, createCoordinate(t, 52.370216, 4.895168), observations{
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 1),
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 2),
				}),
				createCluster(t, createCoordinate(t, 51.9194, 4.4818), observations{
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 3),
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 4),
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 5),
				}),
			},
		},
		{
			name:     "Empty Clusters",
			Clusters: Clusters{},
			want:     Clusters{},
		},
		{
			name: "Cluster with no observations",
			Clusters: Clusters{
				createCluster(t, createCoordinate(t, 52.0, 4.0), observations{
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 1),
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 2),
				}),
				createCluster(t, createCoordinate(t, 51.0, 4.3), observations{}), // empty Cluster
			},
			want: Clusters{
				createCluster(t, createCoordinate(t, 52.370216, 4.895168), observations{
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 1),
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 2),
				}),
				createCluster(t, createCoordinate(t, 51.0, 4.3), observations{}),
			},
		},
	}

	// Loop over test cases and run each one
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := updateCluster(tt.Clusters)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClearObservations(t *testing.T) {
	// Define test cases
	testsTable := []struct {
		name     string
		Clusters Clusters
		want     Clusters
	}{
		{
			name: "Happy path",
			Clusters: Clusters{
				createCluster(t, createCoordinate(t, 52.370216, 4.895168), observations{
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 1),
					createObservation(t, createCoordinate(t, 52.370216, 4.895168), 2),
				}),
				createCluster(t, createCoordinate(t, 51.9194, 4.4818), observations{
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 3),
					createObservation(t, createCoordinate(t, 51.9194, 4.4818), 4),
				}),
			},
			want: Clusters{
				createCluster(t, createCoordinate(t, 52.370216, 4.895168), observations{}),
				createCluster(t, createCoordinate(t, 51.9194, 4.4818), observations{}),
			},
		},
		{
			name:     "Empty Clusters",
			Clusters: Clusters{},
			want:     Clusters{},
		},
	}

	// Run test cases
	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input clusters
			input := make(Clusters, len(tt.Clusters))
			copy(input, tt.Clusters)

			// Call the function
			clearObservations(&input)

			// Check the output
			assert.Equal(t, tt.want, input)
		})
	}
}

func TestInitializeClusters(t *testing.T) {
	// Define some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	obs3 := createObservation(t, createCoordinate(t, 5, 6), 3)
	observations := []observation{obs1, obs2, obs3}

	// Define test cases
	testTable := []struct {
		name                string
		nrClusters          int
		nrCandidateClusters int
		wantError           bool
		expectedError       error
		want                Clusters
	}{
		{
			name:                "NrClusters larger than length of observations",
			nrClusters:          4,
			nrCandidateClusters: 2,
			expectedError:       ErrNrObservationsTooSmall,
			wantError:           true,
			want:                Clusters{},
		},
		{
			name:                "NrCandidateClusters smaller or equal to 0",
			nrClusters:          2,
			nrCandidateClusters: 0,
			expectedError:       ErrNrCandidateClustersTooSmall,
			wantError:           true,
			want:                Clusters{},
		},
		{
			name:                "NrCandidateClusters larger than length of observations",
			nrClusters:          2,
			nrCandidateClusters: 4,
			expectedError:       ErrNrObservationsTooSmall,
			wantError:           true,
			want:                Clusters{},
		},
		{
			name:                "NrClusters equal to length of observations",
			nrClusters:          3,
			nrCandidateClusters: 2,
			expectedError:       nil,
			wantError:           false,
			want: Clusters{
				createCluster(t, obs3.Coordinates, []observation{obs3}),
				createCluster(t, obs1.Coordinates, []observation{obs1}),
				createCluster(t, obs1.Coordinates, []observation{obs1}),
			},
		},
		{
			name:                "Happy path",
			nrClusters:          2,
			nrCandidateClusters: 2,
			expectedError:       nil,
			wantError:           false,
			want: Clusters{
				createCluster(t, obs3.Coordinates, []observation{obs3}),
				createCluster(t, obs1.Coordinates, []observation{obs1}),
			},
		},
	}

	// Loop over test cases and run each one
	// Seed the random number generator to ensure reproducibility
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			randSeed := rand.New(rand.NewSource(12345))
			// Call the function
			Clusters := Clusters{}
			got, err := initializeClusters(observations, Clusters, tt.nrClusters, tt.nrCandidateClusters, randSeed)

			// Check the result
			if tt.wantError {
				assert.Error(t, err, tt.expectedError)
			}
			if !tt.wantError {
				assert.Equal(t, got, tt.want)
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
	want := Clusters{
		createCluster(t, obs3.Coordinates, []observation{obs3}),
		createCluster(t, obs2.Coordinates, []observation{obs2}),
	}

	got := chooseCandidates(observations, probabilities, nrCandidateClusters, randSeed)

	assert.Equal(t, want, got)
}

func TestTotalSumDistance(t *testing.T) {
	// Define some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	obs3 := createObservation(t, createCoordinate(t, 5, 6), 1)
	observations := []observation{obs1, obs2, obs3}

	cluster1 := createCluster(t, obs1.Coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.Coordinates, []observation{obs2})
	Clusters := []Cluster{cluster1, cluster2}

	// Define test cases
	testsTable := []struct {
		name         string
		observations []observation
		checkCluster []Cluster
		want         float64
		wantErr      bool
		err          error
	}{
		{
			name:         "normal case",
			observations: observations,
			checkCluster: Clusters,
			want:         math.Sqrt(8),
			wantErr:      false,
			err:          nil,
		},
		{
			name:         "empty observations",
			observations: []observation{},
			checkCluster: Clusters,
			want:         0,
			wantErr:      false,
			err:          nil,
		},
		{
			name:         "empty Clusters",
			observations: observations,
			checkCluster: []Cluster{},
			want:         0,
			wantErr:      true,
			err:          ErrNoClusters,
		},
	}

	// Loop over test cases and run each one
	for _, tt := range testsTable {
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

	cluster1 := createCluster(t, obs1.Coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.Coordinates, []observation{obs2})
	Clusters := []Cluster{cluster1, cluster2}

	// Define test cases
	tests := []struct {
		name          string
		observations  []observation
		Clusters      []Cluster
		wantDistances []float64
	}{
		{
			name:          "normal case",
			observations:  observations,
			Clusters:      Clusters,
			wantDistances: []float64{math.Sqrt(8), 0, math.Sqrt(8)},
		},
		{
			//since the makes a list based on the length of observations and loops over observations we want to recieve an empty list back
			name:          "empty observations",
			observations:  []observation{},
			Clusters:      Clusters,
			wantDistances: []float64{},
		},
		{
			//because empty Cluster defaults to Center {0,0} we want the distances measured to 0
			name:          "empty Clusters",
			observations:  observations,
			Clusters:      []Cluster{},
			wantDistances: []float64{2.23606797749979, 5, 7.810249675906654},
		},
	}

	// Loop over test cases and run each one
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotDistances []float64
			if len(tt.Clusters) > 0 {
				gotDistances = totalListDistance(tt.observations, tt.Clusters[len(tt.Clusters)-1])
			} else {
				gotDistances = totalListDistance(tt.observations, Cluster{})
			}

			// Check if the computed distances match the expected distances
			assert.Equal(t, tt.wantDistances, gotDistances, "totalListDistance() distances")
		})
	}

}
func TestProbabilities(t *testing.T) {
	// Define some test data
	distances := []float64{1.0, 2.0, 3.0}

	// Call the function and check the result
	want := []float64{1.0, 4.0, 9.0}
	got := Probabilities(distances)
	assert.Equal(t, want, got)
}
func TestBestCandidateFunc(t *testing.T) {
	// Create some test data
	obs1 := createObservation(t, createCoordinate(t, 1, 2), 1)
	obs2 := createObservation(t, createCoordinate(t, 3, 4), 2)
	observations := []observation{obs1, obs2}

	cluster1 := createCluster(t, obs1.Coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.Coordinates, []observation{obs2})
	Clusters := []Cluster{cluster1, cluster2}

	candidate1 := Cluster{Center: Coordinates{Latitude: 5.0, Longitude: 6.0}, observations: []observation{}}
	candidate2 := Cluster{Center: Coordinates{Latitude: 7.0, Longitude: 8.0}, observations: []observation{}}
	candidates := []Cluster{candidate1, candidate2}

	t.Run("Normal run", func(t *testing.T) {
		// Call the function and check the result
		want := candidate1
		got, err := bestCandidateFunc(candidates, Clusters, observations)
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
		candidates := []Cluster{middleCandidate1, middleCandidate2}
		//since the comparison is < we expect the first candidate if the result is equal
		want := middleCandidate1
		// Call the function and check the result
		got, err := bestCandidateFunc(candidates, []Cluster{}, observationsEqual)
		assertEqual(t, got, want, err)
	})
	t.Run("No candidates given", func(t *testing.T) {
		candidates := []Cluster{}
		_, err := bestCandidateFunc(candidates, Clusters, observations)
		assertError(t, err, ErrNrCandidateClustersTooSmall)
	})

}
func TestDistanceToNearestCluster(t *testing.T) {
	// Create Clusters to test with
	Clusters := Clusters{
		createCluster(t, createCoordinate(t, 4, 6), make(observations, 0)),
		createCluster(t, createCoordinate(t, 0, 0), make(observations, 0)),
	}

	// Define test cases as a slice of structs
	testCases := []struct {
		name         string
		observation  observation
		wantDistance float64
		wantIndex    int
	}{
		{
			name:         "Center difference on Longitude",
			observation:  createObservation(t, createCoordinate(t, 2, 6), 1),
			wantDistance: 2.0,
			wantIndex:    0,
		},
		{
			name:         "Center difference on Latitude",
			observation:  createObservation(t, createCoordinate(t, 4, 8), 1),
			wantDistance: 2.0,
			wantIndex:    0,
		},
		{
			name:         "Point is in the middle of two centers",
			observation:  createObservation(t, createCoordinate(t, 2, 3), 1),
			wantDistance: 3.605551275464,
			wantIndex:    0,
		},
		{
			name:         "Point is on Center",
			observation:  createObservation(t, createCoordinate(t, 0, 0), 1),
			wantDistance: 0.0,
			wantIndex:    1,
		},
	}

	// GIS/handheld gps precision
	epsilon := 0.00001

	// Loop over each test case and run the test
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the distanceToNearestCluster function and get the results
			gotDistance, gotIndex, err := distanceToNearestCluster(tc.observation, Clusters)

			// Check for errors
			assert.NoError(t, err, "unexpected error")

			// Check if the distance and Cluster index are correct
			assert.InDeltaf(t, tc.wantDistance, gotDistance, epsilon, "distance: got %f, want %f", gotDistance, tc.wantDistance)
			assert.Equalf(t, tc.wantIndex, gotIndex, "Cluster index: got %d, want %d", gotIndex, tc.wantIndex)
		})
	}
}

func TestDistance(t *testing.T) {

	t.Run("easy run", func(t *testing.T) {
		//create an observation and a Cluster with known Coordinates
		observation := createObservation(t, createCoordinate(t, 10, 20), 1)
		Cluster := createCluster(t, createCoordinate(t, 15, 25), make(observations, 0))
		//calculate the expected distance between the observation and the Cluster
		expectedDistance := math.Sqrt((observation.Coordinates.Latitude-Cluster.Center.Latitude)*(observation.Coordinates.Latitude-Cluster.Center.Latitude) +
			(observation.Coordinates.Longitude-Cluster.Center.Longitude)*(observation.Coordinates.Longitude-Cluster.Center.Longitude))

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, Cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})
	t.Run("Latitude negative", func(t *testing.T) {
		//create an observation and a Cluster with known Coordinates
		observation := createObservation(t, createCoordinate(t, -5, 5), 1)
		Cluster := createCluster(t, createCoordinate(t, 5, 5), make(observations, 0))
		//calculate the expected distance between the observation and the Cluster
		expectedDistance := 10.0

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, Cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})
	t.Run("Longitude negative", func(t *testing.T) {
		//create an observation and a Cluster with known Coordinates
		observation := createObservation(t, createCoordinate(t, 5, -5), 1)
		Cluster := createCluster(t, createCoordinate(t, 5, 5), make(observations, 0))
		//calculate the expected distance between the observation and the Cluster
		expectedDistance := 10.0

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, Cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})
	t.Run("complex float", func(t *testing.T) {
		//create an observation and a Cluster with known Coordinates
		observation := createObservation(t, createCoordinate(t, 51.4486098, 5.4907148), 1)
		Cluster := createCluster(t, createCoordinate(t, 51.395216, 5.474121), make(observations, 0))
		//calculate the expected distance between the observation and the Cluster
		expectedDistance := math.Sqrt((observation.Coordinates.Latitude-Cluster.Center.Latitude)*(observation.Coordinates.Latitude-Cluster.Center.Latitude) +
			(observation.Coordinates.Longitude-Cluster.Center.Longitude)*(observation.Coordinates.Longitude-Cluster.Center.Longitude))

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, Cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0000001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})

	t.Run("distance with 0", func(t *testing.T) {
		//create an observation and a Cluster with known Coordinates
		observation := createObservation(t, createCoordinate(t, 2, 3), 1)
		Cluster := createCluster(t, createCoordinate(t, 0, 0), make(observations, 0))
		//calculate the expected distance between the observation and the Cluster
		expectedDistance := 3.605551275464

		//call the distance function to calculate the actual distance
		actualDistance := distance(observation, Cluster)

		//check if the actual distance matches the expected distance
		if !AlmostEqual(t, actualDistance, expectedDistance, 0.0000001) {
			t.Errorf("distance calculation incorrect, expected: %v, got: %v", expectedDistance, actualDistance)
		}
	})

}

func TestActivitiesToObservations(t *testing.T) {
	activity1 := makeActivity(t, "1", "51.5074", "-0.1278", "1234")
	activity2 := makeActivity(t, "2", "52.5200", "13.4050", "1234")
	activities := activities{
		*activity1,
		*activity2,
	}
	want := observations{
		{
			id: 1,
			Coordinates: Coordinates{
				Latitude:  51.5074,
				Longitude: -0.1278,
			},
		},
		{
			id: 2,
			Coordinates: Coordinates{
				Latitude:  52.5200,
				Longitude: 13.4050,
			},
		},
	}

	var observations = make(observations, 0)
	got, err := activitiesToObservations(activities, observations)

	assertEqual(t, got, want, err)

}

func TestHaversineDistance(t *testing.T) {
	// Define some test data
	point1 := Coordinates{Latitude: 51.4416, Longitude: 5.4697} // btw Eindhoven :D
	point2 := Coordinates{Latitude: 48.8566, Longitude: 2.3522} // OUI Paris

	// Call the function
	got := haversineDistance(point1, point2)

	// Check the result
	expected := 363.2 // expected distance in kilometers
	if !AlmostEqual(t, got, expected, 0.01) {
		t.Errorf("haversineDistance() returned %v km, want %v km", got, expected)
	}
}

func TestDistanceKilometers(t *testing.T) {
	testCases := []struct {
		name           string
		observation    observation
		cluster        Cluster
		expectedResult float64
	}{
		{
			name: "Basic test case",
			observation: observation{
				Coordinates: createCoordinate(t, 40.7128, -74.0060),
			},
			cluster: Cluster{
				Center: createCoordinate(t, 37.7749, -122.4194),
			},
			expectedResult: 4129,
		},
		{
			name: "Observation and cluster have same coordinates",
			observation: observation{
				Coordinates: createCoordinate(t, 37.7749, -122.4194),
			},
			cluster: Cluster{
				Center: createCoordinate(t, 37.7749, -122.4194),
			},
			expectedResult: 0,
		},
		{
			name: "Observation and cluster are on opposite sides of the earth",
			observation: observation{
				Coordinates: createCoordinate(t, 51.448557, 5.450123),
			},
			cluster: Cluster{
				Center: createCoordinate(t, -51.448557, -174.549877),
			},
			expectedResult: 20015,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := distanceKilometers(tt.observation, tt.cluster)
			assert.InDelta(t, tt.expectedResult, result, 1.0, "Unexpected distance result")
		})
	}
}

func BenchmarkKMeansALot(t *testing.B) {
	activities := make(activities, 0)
	for i := 0; i < 100000; i++ {
		str1 := strconv.Itoa(rand.Intn(1000))
		str2 := strconv.Itoa(rand.Intn(1000))
		iStr := strconv.Itoa(i)
		result := makeActivity(t, iStr, str1, str2, "1234")
		activities = append(activities, *result)
	}
	KMeans(activities, 30, 10)
}

func makeActivity(t testing.TB, id string, lat string, long string, zipcode string) *model.ActivityModelBumbal {
	t.Helper()

	//make activity and addressapplied model
	activity := model.NewActivityModel()
	address := model.NewAddressAppliedModel()
	//set lat and long
	address.SetLatitude(lat)
	address.SetLongitude(long)
	address.SetZipcode(zipcode)
	activity.SetAddressApplied(*address)
	activity.SetId(id)

	return activity
}

func createCluster(t testing.TB, Center Coordinates, observations observations) Cluster {
	t.Helper()
	Cluster := Cluster{
		Center:       Center,
		observations: observations,
	}
	return Cluster
}

func createObservation(t testing.TB, Coordinates Coordinates, id int64) observation {
	t.Helper()
	observation := observation{
		id:          id,
		Coordinates: Coordinates,
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

func createCoordinate(t testing.TB, lat float64, long float64) Coordinates {
	t.Helper()
	coord := Coordinates{
		Latitude:  lat,
		Longitude: long,
	}

	return coord
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
