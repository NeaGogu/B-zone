package kMeans

import (
	"bzone/backend/internal/models"
	model "bzone/backend/internal/models"
	"errors"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MinNormal is the smallest positive normal value of type float64.
var MinNormal = math.Float64frombits(0x0010000000000000)

// func TestKMeans(t *testing.T) {
// 	// Define some test data
// 	activity1 := makeActivity(t, "1", "1", "1", "1234")
// 	activity2 := makeActivity(t, "1", "3", "4", "1234")
// 	activity3 := makeActivity(t, "1", "5", "6", "1234")
// 	activities := []models.ActivityModelBumbal{*activity1, *activity2, *activity3}

// 	testCases := []struct {
// 		name                string
// 		activities          []models.ActivityModelBumbal
// 		nrClusters          int
// 		nrCandidateClusters int
// 		expectedError       error
// 	}{
// 		{
// 			name:                "NrClusters smaller than 1",
// 			activities:          activities,
// 			nrClusters:          0,
// 			nrCandidateClusters: 2,
// 			expectedError:       ErrNrClustersTooSmall,
// 		},
// 		{
// 			name:                "NrCandidateClusters smaller or equal to 0",
// 			activities:          activities,
// 			nrClusters:          2,
// 			nrCandidateClusters: 0,
// 			expectedError:       ErrNrCandidateClustersTooSmall,
// 		},
// 		{
// 			name:                "Happy path",
// 			activities:          activities,
// 			nrClusters:          2,
// 			nrCandidateClusters: 2,
// 			expectedError:       nil,
// 		},
// 		{
// 			name:                "No activities",
// 			activities:          []models.ActivityModelBumbal{},
// 			nrClusters:          2,
// 			nrCandidateClusters: 2,
// 			expectedError:       ErrNoActivities,
// 		},
// 	}

// 	// Loop over test cases and run each one
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Call the function
// 			got, err := KMeans(tc.activities, tc.nrClusters, tc.nrCandidateClusters)

// 			// Check the result
// 			if err != tc.expectedError {
// 				t.Errorf("kMeans() error = %v, want %v", err, tc.expectedError)
// 			}

// 			// If there was no error, check that the number of Clusters matches the input
// 			if err == nil && len(got) != tc.nrClusters {
// 				t.Errorf("kMeans() returned %v Clusters, want %v Clusters", len(got), tc.nrClusters)
// 			}

// 			// Check that each Cluster has at least one observation
// 			for _, c := range got {
// 				if len(c.ZoneRanges) == 0 {
// 					t.Errorf("kMeans() returned a Cluster with 0 observations")
// 				}
// 			}
// 		})
// 	}
// }

func TestUpdateCluster(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		Clusters Clusters
		want     Clusters
		wantErr  bool
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
			wantErr: false,
		},
		{
			name:     "Empty Clusters",
			Clusters: Clusters{},
			want:     Clusters{},
			wantErr:  false,
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
			wantErr: false,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := updateCluster(tt.Clusters)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClearObservations(t *testing.T) {
	// Define test cases
	tests := []struct {
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

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a deep copy of the input Clusters
			input := make(Clusters, len(tt.Clusters))
			copy(input, tt.Clusters)

			// Call the function
			clearObservations(&input)

			// Check if the output matches the expected output
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("clearObservations() got %v, want %v", input, tt.want)
			}
		})
	}
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
		expectedClusters    Clusters
	}{
		{
			name:                "NrClusters larger than length of observations",
			nrClusters:          4,
			nrCandidateClusters: 2,
			expectedError:       ErrNrObservationsTooSmall,
			wantError:           true,
			expectedClusters:    Clusters{},
		},
		{
			name:                "NrCandidateClusters smaller or equal to 0",
			nrClusters:          2,
			nrCandidateClusters: 0,
			expectedError:       ErrNrCandidateClustersTooSmall,
			wantError:           true,
			expectedClusters:    Clusters{},
		},
		{
			name:                "NrCandidateClusters larger than length of observations",
			nrClusters:          2,
			nrCandidateClusters: 4,
			expectedError:       ErrNrObservationsTooSmall,
			wantError:           true,
			expectedClusters:    Clusters{},
		},
		{
			name:                "NrClusters equal to length of observations",
			nrClusters:          3,
			nrCandidateClusters: 2,
			expectedError:       nil,
			wantError:           false,
			expectedClusters: Clusters{
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
			expectedClusters: Clusters{
				createCluster(t, obs3.Coordinates, []observation{obs3}),
				createCluster(t, obs1.Coordinates, []observation{obs1}),
			},
		},
	}

	// Loop over test cases and run each one
	// Seed the random number generator to ensure reproducibility
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			randSeed := rand.New(rand.NewSource(12345))
			// Call the function
			Clusters := Clusters{}
			got, err := initializeClusters(observations, Clusters, tc.nrClusters, tc.nrCandidateClusters, randSeed)

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

	// Define a test table with different scenarios
	tests := []struct {
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

	cluster1 := createCluster(t, obs1.Coordinates, []observation{obs1})
	cluster2 := createCluster(t, obs2.Coordinates, []observation{obs2})
	Clusters := []Cluster{cluster1, cluster2}

	// Define a test table with different scenarios
	tests := []struct {
		name          string
		observations  []observation
		Clusters      []Cluster
		wantDistances []float64
		wantErr       bool
		err           error
	}{
		{
			name:          "normal case",
			observations:  observations,
			Clusters:      Clusters,
			wantDistances: []float64{math.Sqrt(8), 0, math.Sqrt(8)},
			wantErr:       false,
			err:           nil,
		},
		{
			//since the makes a list based on the length of observations and loops over observations we want to recieve an empty list back
			name:          "empty observations",
			observations:  []observation{},
			Clusters:      Clusters,
			wantDistances: []float64{},
			wantErr:       false,
			err:           nil,
		},
		{
			//because empty Cluster defaults to Center {0,0} we want the distances measured to 0
			name:          "empty Clusters",
			observations:  observations,
			Clusters:      []Cluster{},
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
			if len(tt.Clusters) > 0 {
				gotDistances = totalListDistance(tt.observations, tt.Clusters[len(tt.Clusters)-1])
			} else {
				gotDistances = totalListDistance(tt.observations, Cluster{})
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

func TestClusterToZoneModel(t *testing.T) {
	// create sample data
	activity1 := makeActivity(t, "1", "52.3764", "4.9004", "1000ab")
	activity2 := makeActivity(t, "2", "52.3764", "4.9004", "1000cd")
	activity3 := makeActivity(t, "3", "51.2194", "4.4025", "2000da")
	activity4 := makeActivity(t, "4", "51.2194", "4.4025", "2000ab")

	tests := []struct {
		name       string
		Clusters   Clusters
		activities []model.ActivityModelBumbal
		want       []model.ZoneModel
		wantErr    bool
	}{
		{
			name: "happy path",
			Clusters: Clusters{
				Cluster{Center: createCoordinate(t, 1, 2),
					observations: observations{createObservation(t, createCoordinate(t, 1, 2), 1), createObservation(t, createCoordinate(t, 1, 2), 2)}},
				Cluster{Center: createCoordinate(t, 1, 2),
					observations: observations{createObservation(t, createCoordinate(t, 1, 2), 3), createObservation(t, createCoordinate(t, 1, 2), 4)}},
			},
			activities: []model.ActivityModelBumbal{
				*activity1,
				*activity2,
				*activity3,
				*activity4,
			},
			want: []model.ZoneModel{
				{
					Id:         "0",
					ZoneRanges: []model.ZoneRangeModel{{ZoneRangeId: 0, ZipcodeFrom: 1000, ZipcodeTo: 1000, IsoCountry: "NLD"}},
				},
				{
					Id:         "1",
					ZoneRanges: []model.ZoneRangeModel{{ZoneRangeId: 1, ZipcodeFrom: 2000, ZipcodeTo: 2000, IsoCountry: "NLD"}},
				},
			},
			wantErr: false,
		},
		{
			name:     "empty Clusters",
			Clusters: Clusters{},
			activities: []model.ActivityModelBumbal{
				*activity1,
				*activity2,
				*activity3,
				*activity4,
			},
			want:    []model.ZoneModel{},
			wantErr: false,
		},
		{
			name: "empty activities",
			Clusters: Clusters{
				Cluster{Center: createCoordinate(t, 1, 2),
					observations: observations{createObservation(t, createCoordinate(t, 1, 2), 1), createObservation(t, createCoordinate(t, 1, 2), 2)}},
				Cluster{Center: createCoordinate(t, 1, 2),
					observations: observations{createObservation(t, createCoordinate(t, 1, 2), 3), createObservation(t, createCoordinate(t, 1, 2), 4)}},
			},
			activities: []model.ActivityModelBumbal{},
			want:       nil,
			wantErr:    true,
		},
		{
			name: "invalid activity id",
			Clusters: Clusters{
				Cluster{Center: createCoordinate(t, 1, 2),
					observations: observations{createObservation(t, createCoordinate(t, 1, 2), 1), createObservation(t, createCoordinate(t, 1, 2), 2)}},
				Cluster{Center: createCoordinate(t, 1, 2),
					observations: observations{createObservation(t, createCoordinate(t, 1, 2), 3), createObservation(t, createCoordinate(t, 1, 2), 4)}},
			},
			activities: []model.ActivityModelBumbal{
				*activity1,
				*activity2,
				*activity3,
				*makeActivity(t, "invalid", "1", "1", "5923"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := clusterToZoneModel(tt.Clusters, tt.activities)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			for _, zone := range tt.want {
				assert.Contains(t, got, zone)
			}
			assert.Equal(t, len(tt.want), len(got))
		})
	}
}
func TestClusterToZipcodeSet(t *testing.T) {
	// create sample data
	activity1 := makeActivity(t, "1", "52.3764", "4.9004", "1012")
	activity2 := makeActivity(t, "2", "52.3764", "4.9004", "1017")
	activity3 := makeActivity(t, "3", "51.2194", "4.4025", "2273")
	cluster1 := createCluster(t, Coordinates{52.3764, 4.9004}, observations{observation{id: 1, Coordinates: createCoordinate(t, 4.9004, 52.3764)}, observation{id: 2, Coordinates: createCoordinate(t, 4.9004, 52.3764)}})
	cluster2 := createCluster(t, Coordinates{51.2194, 4.4025}, observations{observation{id: 3, Coordinates: createCoordinate(t, 4.4025, 51.2194)}})
	Clusters := Clusters{cluster1, cluster2}
	activities := []models.ActivityModelBumbal{*activity1, *activity2, *activity3}

	// call the function
	result, err := clusterToZipcodeSet(Clusters, activities)
	if err != nil {
		t.Errorf("Didn't expect an error, but got %q", err)
	}

	// assert the result
	expected := []map[string]struct{}{
		{"1012": {}, "1017": {}},
		{"2273": {}},
	}

	assert.Equal(t, expected, result)
}

func TestZipcodeSetToZoneModel(t *testing.T) {
	tests := []struct {
		name          string
		listZipcodes  []map[string]struct{}
		expectedZones []model.ZoneModel
		expectedError error
	}{
		{
			name:          "Empty list of zip code sets",
			listZipcodes:  []map[string]struct{}{},
			expectedZones: []model.ZoneModel{},
			expectedError: nil,
		},
		{
			name: "Single zip code set",
			listZipcodes: []map[string]struct{}{
				{"1234": {}, "3456": {}},
			},
			expectedZones: []model.ZoneModel{
				{
					Id: "0",
					ZoneRanges: []model.ZoneRangeModel{
						{ZoneRangeId: 0, ZipcodeFrom: 1234, ZipcodeTo: 1234, IsoCountry: "NLD"},
						{ZoneRangeId: 1, ZipcodeFrom: 3456, ZipcodeTo: 3456, IsoCountry: "NLD"},
					},
					ZoneFuelCost:    0,
					ZoneDrivingTime: 0,
				},
			},
			expectedError: nil,
		},
		{
			name: "Multiple zip code sets",
			listZipcodes: []map[string]struct{}{
				{"1234": {}, "2345": {}},
				{"3456": {}, "4567": {}},
			},
			expectedZones: []model.ZoneModel{
				{
					Id: "0",
					ZoneRanges: []model.ZoneRangeModel{
						{ZoneRangeId: 0, ZipcodeFrom: 1234, ZipcodeTo: 1234, IsoCountry: "NLD"},
						{ZoneRangeId: 1, ZipcodeFrom: 2345, ZipcodeTo: 2345, IsoCountry: "NLD"},
					},
					ZoneFuelCost:    0,
					ZoneDrivingTime: 0,
				},
				{
					Id: "1",
					ZoneRanges: []model.ZoneRangeModel{
						{ZoneRangeId: 1, ZipcodeFrom: 3456, ZipcodeTo: 3456, IsoCountry: "NLD"},
						{ZoneRangeId: 2, ZipcodeFrom: 4567, ZipcodeTo: 4567, IsoCountry: "NLD"},
					},
					ZoneFuelCost:    0,
					ZoneDrivingTime: 0,
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid zip code set (too short)",
			listZipcodes: []map[string]struct{}{
				{"123a": {}, "2345": {}},
			},
			expectedZones: nil,
			expectedError: errors.New("zipcode \"123a\" trimmed to \"123\" is not 4 digits long"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualZones, actualError := zipcodeSetToZoneModel(tt.listZipcodes)
			assert.Equal(t, tt.expectedError, actualError)
			assert.Equal(t, tt.expectedZones, actualZones)
		})
	}
}

func TestCreateZoneRanges(t *testing.T) {
	tests := []struct {
		name                 string
		idCounter            int
		zipcodeSet           map[string]struct{}
		expectedZoneRanges   []model.ZoneRangeModel
		expectedErrorMessage string
	}{
		{
			name:                 "Empty zip code set",
			idCounter:            0,
			zipcodeSet:           map[string]struct{}{},
			expectedZoneRanges:   nil,
			expectedErrorMessage: "length of zipcodeSet is too small: map[]",
		},
		{
			name:       "Single zip code set",
			idCounter:  0,
			zipcodeSet: map[string]struct{}{"1234": {}, "1235": {}, "1236": {}},
			expectedZoneRanges: []model.ZoneRangeModel{
				{
					ZoneRangeId: 0,
					ZipcodeFrom: 1234,
					ZipcodeTo:   1236,
					IsoCountry:  "NLD",
				},
			},
			expectedErrorMessage: "",
		},
		{
			name:       "Multiple zip code sets",
			idCounter:  0,
			zipcodeSet: map[string]struct{}{"1234": {}, "1235": {}, "2345": {}},
			expectedZoneRanges: []model.ZoneRangeModel{
				{
					ZoneRangeId: 0,
					ZipcodeFrom: 1234,
					ZipcodeTo:   1235,
					IsoCountry:  "NLD",
				},
				{
					ZoneRangeId: 1,
					ZipcodeFrom: 2345,
					ZipcodeTo:   2345,
					IsoCountry:  "NLD",
				},
			},
			expectedErrorMessage: "",
		},
		{
			name:                 "Invalid zip code",
			idCounter:            0,
			zipcodeSet:           map[string]struct{}{"123b": {}, "2345": {}},
			expectedZoneRanges:   nil,
			expectedErrorMessage: "zipcode \"123b\" trimmed to \"123\" is not 4 digits long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualZoneRanges, actualError := createZoneRanges(tt.idCounter, tt.zipcodeSet)
			if actualError == nil && tt.expectedErrorMessage != "" {
				t.Errorf("expected error message '%v', but got nil", tt.expectedErrorMessage)
			} else if actualError != nil && tt.expectedErrorMessage == "" {
				t.Errorf("expected no error, but got '%v'", actualError)
			} else if actualError != nil && tt.expectedErrorMessage != actualError.Error() {
				t.Errorf("expected error message '%v', but got '%v'", tt.expectedErrorMessage, actualError.Error())
			}
			assert.Equal(t, tt.expectedZoneRanges, actualZoneRanges)
		})
	}
}

func TestCreateZoneRange(t *testing.T) {
	tests := []struct {
		name                 string
		id                   string
		zipcodeFrom          int64
		zipcodeTo            int64
		isoCountry           string
		expectedZoneRange    model.ZoneRangeModel
		expectedErrorMessage string
	}{
		{
			name:        "Valid case",
			id:          "1",
			zipcodeFrom: 1000,
			zipcodeTo:   2000,
			isoCountry:  "NLD",
			expectedZoneRange: model.ZoneRangeModel{
				ZoneRangeId: 1,
				ZipcodeFrom: 1000,
				ZipcodeTo:   2000,
				IsoCountry:  "NLD",
			},
			expectedErrorMessage: "",
		},
		{
			name:                 "Invalid case: invalid zipcode range",
			id:                   "1",
			zipcodeFrom:          2000,
			zipcodeTo:            1000,
			isoCountry:           "NLD",
			expectedZoneRange:    model.ZoneRangeModel{},
			expectedErrorMessage: "zipcodeFrom cannot be greater than zipcodeTo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualZoneRange, actualError := createZoneRange(tt.id, tt.zipcodeFrom, tt.zipcodeTo, tt.isoCountry)
			if actualError == nil && tt.expectedErrorMessage != "" {
				t.Errorf("expected error message '%v', but got nil", tt.expectedErrorMessage)
			} else if actualError != nil && tt.expectedErrorMessage == "" {
				t.Errorf("expected no error, but got '%v'", actualError)
			} else if actualError != nil && tt.expectedErrorMessage != actualError.Error() {
				t.Errorf("expected error message '%v', but got '%v'", tt.expectedErrorMessage, actualError.Error())
			}
			if !reflect.DeepEqual(actualZoneRange, tt.expectedZoneRange) {
				t.Errorf("unexpected zone range: got %v, want %v", actualZoneRange, tt.expectedZoneRange)
			}
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

func makeActivity(t testing.TB, id string, lat string, long string, zipcode string) *models.ActivityModelBumbal {
	t.Helper()

	//make activity and addressapplied model
	activity := models.NewActivityModel()
	address := models.NewAddressAppliedModel()
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

func createObservations(t testing.TB, coordinatesList []Coordinates) observations {
	t.Helper()

	observations := make(observations, 0)

	for i, coord := range coordinatesList {
		observations = append(observations, observation{
			id:          int64(i),
			Coordinates: coord,
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
