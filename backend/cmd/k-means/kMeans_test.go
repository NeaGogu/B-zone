package kMeans

import (
	"testing"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func TestinitCentroids(t *testing.T) {

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
