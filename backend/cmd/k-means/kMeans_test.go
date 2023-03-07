package kMeans

import (
	"testing"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func TestinitCentroids(t *testing.T) {

	t.Run("len activities less than or equal to 0", func(t *testing.T) {
		activities := []openapi.ActivityModel{
			{},
			{},
		}
		nrClusters := 1
		nrCandidateCenters := 1
		_, got := initCentroids(activities, nrClusters, nrCandidateCenters)
		want := ErrNoActivities

		if got == nil {
			t.Errorf("didn't get an error but wanted %q", want)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("nrClusters less than activities", func(t *testing.T) {
		activities := []openapi.ActivityModel{
			{},
			{},
		}
		nrClusters := 1
		nrCandidateCenters := 1
		_, got := initCentroids(activities, nrClusters, nrCandidateCenters)
		want := ErrNrActivitiesTooSmall

		if got == nil {
			t.Errorf("didn't get an error but wanted %q", want)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("nrCandidateCenters less than or equal to 0", func(t *testing.T) {
		activities := []openapi.ActivityModel{
			{},
			{},
		}
		nrClusters := 1
		nrCandidateCenters := 0
		_, got := initCentroids(activities, nrClusters, nrCandidateCenters)
		want := ErrNrCandidateCentersTooSmall

		if got == nil {
			t.Errorf("didn't get an error but wanted %q", want)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
