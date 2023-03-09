package kMeans

import (
	"reflect"
	"strconv"
	"testing"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/mroth/weightedrand/v2"
)

func TestInitCentroids(t *testing.T) {

	activitiesEmpty := []openapi.ActivityModel{}
	activitiesTwo := []openapi.ActivityModel{
		{},
		{},
	}

	t.Run("len activities less than or equal to 0", func(t *testing.T) {
		nrClusters := 1
		nrCandidateCenters := 1
		_, got := initCentroids(activitiesEmpty, nrClusters, nrCandidateCenters)
		want := ErrNoActivities

		assertError(t, got, want)
	})

	t.Run("nrClusters larger than number of activities", func(t *testing.T) {
		nrClusters := 3
		nrCandidateCenters := 1
		_, got := initCentroids(activitiesTwo, nrClusters, nrCandidateCenters)
		want := ErrNrActivitiesTooSmall

		assertError(t, got, want)
	})

	t.Run("nrCandidateCenters less than or equal to 0", func(t *testing.T) {
		nrClusters := 2
		nrCandidateCenters := 0
		_, got := initCentroids(activitiesTwo, nrClusters, nrCandidateCenters)
		want := ErrNrCandidateCentersTooSmall

		assertError(t, got, want)
	})

	t.Run("Valid run one activity", func(t *testing.T) {
		activity := makeActivity(t, "1", "1")

		nrClusters := 1
		nrCandidateCenters := 1
		activities := []openapi.ActivityModel{
			*activity,
		}
		got, err := initCentroids(activities, nrClusters, nrCandidateCenters)

		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}

		want := []centroid{
			makeCentroid(t, 1, 1, []openapi.ActivityModel{*activity}),
		}

		assertEqual(t, got, want, err)
	})

	t.Run("Valid run multiple activity", func(t *testing.T) {

		activity := makeActivity(t, "1", "1")
		activity2 := makeActivity(t, "1", "2")
		nrClusters := 2
		nrCandidateCenters := 2
		activities := []openapi.ActivityModel{
			*activity,
			*activity2,
		}
		got, err := initCentroids(activities, nrClusters, nrCandidateCenters)

		want := []centroid{
			makeCentroid(t, 1, 1, []openapi.ActivityModel{*activity}),
			makeCentroid(t, 1, 2, []openapi.ActivityModel{*activity2}),
		}

		assertEqual(t, got, want, err)
	})

}

func TestSampleCandidateCentroid(t *testing.T) {

	//Create first activity
	activity1 := makeActivity(t, "1", "1")
	//create second activity
	activity2 := makeActivity(t, "2", "2")
	//assign activities to slice
	activities := []openapi.ActivityModel{
		*activity1,
		*activity2,
	}

	t.Run("valid run empty candidateCenter", func(t *testing.T) {
		randomInt := 1
		candidateCenter := []centroid{}
		got, err := sampleCandidateCentroid(activities, randomInt, candidateCenter)

		want := []centroid{
			{
				center: coordinates{
					lat:  2.00,
					long: 2.00,
				},
				AssignedActivities: []openapi.ActivityModel{*activity2},
			},
		}

		assertEqual(t, got, want, err)
	})

	t.Run("valid run non empty candidateCenter", func(t *testing.T) {

		randomInt := 1
		candidateCenter := []centroid{
			{
				center: coordinates{
					lat:  1.00,
					long: 1.00,
				},
				AssignedActivities: []openapi.ActivityModel{*activity1},
			},
		}
		got, err := sampleCandidateCentroid(activities, randomInt, candidateCenter)

		want := []centroid{
			{
				center: coordinates{
					lat:  1.00,
					long: 1.00,
				},
				AssignedActivities: []openapi.ActivityModel{*activity1},
			},
			{
				center: coordinates{
					lat:  2.00,
					long: 2.00,
				},
				AssignedActivities: []openapi.ActivityModel{*activity2},
			},
		}
		assertEqual(t, got, want, err)
	})

	t.Run("invalid float lat", func(t *testing.T) {
		activity := makeActivity(t, "a.00", "1.00")
		activities := []openapi.ActivityModel{
			*activity,
		}

		_, got := sampleCandidateCentroid(activities, 0, emptyCentroid)
		e := got.(*strconv.NumError)
		want := strconv.ErrSyntax

		assertError(t, e.Err, want)

	})

	t.Run("invalid float long", func(t *testing.T) {
		activity := makeActivity(t, "1.00", "a.00")
		activities := []openapi.ActivityModel{
			*activity,
		}

		_, got := sampleCandidateCentroid(activities, 0, emptyCentroid)
		e := got.(*strconv.NumError)
		want := strconv.ErrSyntax

		assertError(t, e.Err, want)
	})

}

func TestNewChoices(t *testing.T) {

	//Create first activity
	activity1 := makeActivity(t, "1", "1")
	//create second activity
	activity2 := makeActivity(t, "2", "2")
	//assign activities to slice
	activities := []openapi.ActivityModel{
		*activity1,
		*activity2,
	}

	t.Run("Valid run", func(t *testing.T) {
		got := newChoices(activities, centroid{
			center: coordinates{
				lat:  1.00,
				long: 1.00,
			},
			AssignedActivities: []openapi.ActivityModel{*activity1},
		})

		want := []weightedrand.Choice[int, int]{
			{Item: 0,
				Weight: 0 * 1000},
			{Item: 1,
				Weight: 1 * 1000},
		}
		assertEqual(t, got, want, nil)
	})
}

func TestKmeans(t *testing.T) {

	//Create first activity
	activity1 := makeActivity(t, "1", "1")
	//create second activity
	activity2 := makeActivity(t, "1", "1")
	activity3 := makeActivity(t, "8", "8")
	activity4 := makeActivity(t, "8", "8")
	//assign activities to slice
	activities := []openapi.ActivityModel{
		*activity1,
		*activity2,
		*activity3,
		*activity4,
	}
	wantedCentroid1 := makeCentroid(t, 1.00, 1.00, []openapi.ActivityModel{
		*activity1, *activity2,
	})
	wantedCentroid2 := makeCentroid(t, 8.00, 8.00, []openapi.ActivityModel{
		*activity3, *activity3,
	})

	want := []centroid{
		wantedCentroid1,
		wantedCentroid2,
	}

	got, err := kMeans(activities, 2)

	assertEqual(t, got, want, err)
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

func makeActivity(t testing.TB, lat string, long string) *openapi.ActivityModel {
	t.Helper()

	//make activity and addressapplied model
	activity := openapi.NewActivityModel()
	address := openapi.NewAddressAppliedModel()
	//set lat and long
	address.SetLatitude(lat)
	address.SetLongitude(long)
	activity.SetAddressApplied(*address)

	return activity
}

func makeCentroid(t testing.TB, lat float64, long float64, assignedActivities []openapi.ActivityModel) centroid {
	t.Helper()
	centroid := centroid{
		center: coordinates{
			lat:  lat,
			long: long,
		},
		AssignedActivities: assignedActivities,
	}
	return centroid
}
