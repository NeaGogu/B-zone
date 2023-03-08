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
		activity := openapi.NewActivityModel()
		address := openapi.NewAddressAppliedModel()
		address.SetLatitude("1.00")
		address.SetLongitude("1.00")
		activity.SetAddressApplied(*address)

		nrClusters := 1
		nrCandidateCenters := 1
		activities := []openapi.ActivityModel{
			*activity,
		}
		//fmt.Print(activities)
		got, err := initCentroids(activities, nrClusters, nrCandidateCenters)

		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}
		want := []centroid{
			{
				center: coordinates{
					long: 1,
					lat:  1,
				},
				AssignedActivities: *activity,
			}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n want %v", got, want)
		}
	})

	t.Run("Valid run multiple activity", func(t *testing.T) {
		activity := openapi.NewActivityModel()
		address := openapi.NewAddressAppliedModel()
		address.SetLatitude("1.00")
		address.SetLongitude("1.00")
		activity.SetAddressApplied(*address)

		nrClusters := 2
		nrCandidateCenters := 2
		activities := []openapi.ActivityModel{
			*activity,
			*activity,
		}
		//fmt.Print(activities)
		got, err := initCentroids(activities, nrClusters, nrCandidateCenters)

		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}
		want := []centroid{
			{
				center: coordinates{
					long: 1,
					lat:  1,
				},
				AssignedActivities: *activity,
			},
			{
				center: coordinates{
					long: 1,
					lat:  1,
				},
				AssignedActivities: *activity,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n want %v", got, want)
		}
	})

}

func TestSampleCandidateCentroid(t *testing.T) {

	//Create first activity
	activity1 := openapi.NewActivityModel()
	address1 := openapi.NewAddressAppliedModel()
	address1.SetLatitude("1.00")
	address1.SetLongitude("1.00")
	activity1.SetAddressApplied(*address1)
	//create second activity
	activity2 := openapi.NewActivityModel()
	address2 := openapi.NewAddressAppliedModel()
	address2.SetLatitude("2.00")
	address2.SetLongitude("2.00")
	activity2.SetAddressApplied(*address2)
	//assign activities to slice
	activities := []openapi.ActivityModel{
		*activity1,
		*activity2,
	}

	t.Run("valid run empty candidateCenter", func(t *testing.T) {
		randomInt := 1
		candidateCenter := []centroid{}
		got, err := sampleCandidateCentroid(activities, randomInt, candidateCenter)

		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}

		want := []centroid{
			{
				center: coordinates{
					lat:  2.00,
					long: 2.00,
				},
				AssignedActivities: *activity2,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n want %v", got, want)
		}
	})

	t.Run("valid run non empty candidateCenter", func(t *testing.T) {

		randomInt := 1
		candidateCenter := []centroid{
			{
				center: coordinates{
					lat:  1.00,
					long: 1.00,
				},
				AssignedActivities: *activity1,
			},
		}
		got, err := sampleCandidateCentroid(activities, randomInt, candidateCenter)

		if err != nil {
			t.Errorf("didn't want an error but got: %q", err)
		}

		want := []centroid{
			{
				center: coordinates{
					lat:  1.00,
					long: 1.00,
				},
				AssignedActivities: *activity1,
			},
			{
				center: coordinates{
					lat:  2.00,
					long: 2.00,
				},
				AssignedActivities: *activity2,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n want %v", got, want)
		}
	})

	t.Run("invalid float lat", func(t *testing.T) {
		activity := openapi.NewActivityModel()
		address := openapi.NewAddressAppliedModel()
		address.SetLatitude("a.00")
		address.SetLongitude("1.00")
		activity.SetAddressApplied(*address)
		activities := []openapi.ActivityModel{
			*activity,
		}

		_, got := sampleCandidateCentroid(activities, 0, emptyCentroid)
		e := got.(*strconv.NumError)
		want := strconv.ErrSyntax

		assertError(t, e.Err, want)

	})

	t.Run("invalid float long", func(t *testing.T) {
		activity := openapi.NewActivityModel()
		address := openapi.NewAddressAppliedModel()
		address.SetLatitude("1.00")
		address.SetLongitude("a.00")
		activity.SetAddressApplied(*address)
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
	activity1 := openapi.NewActivityModel()
	address1 := openapi.NewAddressAppliedModel()
	address1.SetLatitude("1.00")
	address1.SetLongitude("1.00")
	activity1.SetAddressApplied(*address1)
	//create second activity
	activity2 := openapi.NewActivityModel()
	address2 := openapi.NewAddressAppliedModel()
	address2.SetLatitude("2.00")
	address2.SetLongitude("2.00")
	activity2.SetAddressApplied(*address2)
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
			AssignedActivities: *activity1,
		})

		want := []weightedrand.Choice[int, int]{
			{Item: 0,
				Weight: 5 * 100},
			{Item: 1,
				Weight: 5 * 100},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n want %v", got, want)
		}
	})
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
