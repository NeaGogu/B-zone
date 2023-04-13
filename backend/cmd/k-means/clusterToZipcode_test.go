package kMeans

import (
	model "bzone/backend/internal/models"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			got, err := ClusterToZoneModel(tt.Clusters, tt.activities)
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
	activities := []model.ActivityModelBumbal{*activity1, *activity2, *activity3}

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
		expectError          bool
		expectedErrorMessage string
	}{
		{
			name:                 "Empty zip code set",
			idCounter:            0,
			zipcodeSet:           map[string]struct{}{},
			expectedZoneRanges:   nil,
			expectError:          true,
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
			expectError:          false,
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
			expectError:          false,
			expectedErrorMessage: "",
		},
		{
			name:                 "Invalid zip code",
			idCounter:            0,
			zipcodeSet:           map[string]struct{}{"123b": {}, "2345": {}},
			expectedZoneRanges:   nil,
			expectError:          true,
			expectedErrorMessage: "zipcode \"123b\" trimmed to \"123\" is not 4 digits long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualZoneRanges, actualError := createZoneRanges(tt.idCounter, tt.zipcodeSet)

			// Check that the error matches the expected error
			if tt.expectError {
				assert.Error(t, actualError)
				assert.EqualError(t, actualError, tt.expectedErrorMessage)
			} else {
				assert.NoError(t, actualError)
			}

			// Check that the result matches the expected result
			assert.Equal(t, tt.expectedZoneRanges, actualZoneRanges)
		})
	}
}

func TestCreateZoneRange(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		zipcodeFrom    int64
		zipcodeTo      int64
		isoCountry     string
		expectedResult model.ZoneRangeModel
		expectedError  error
	}{
		{
			name:        "Valid range in Netherlands",
			id:          "123",
			zipcodeFrom: 1000, // Example zip code from the Netherlands
			zipcodeTo:   1999, // Example zip code from the Netherlands
			isoCountry:  "NL", // ISO code for the Netherlands
			expectedResult: model.ZoneRangeModel{
				ZoneRangeId: 123,
				ZipcodeFrom: 1000,
				ZipcodeTo:   1999,
				IsoCountry:  "NL",
			},
			expectedError: nil,
		},
		{
			name:           "Invalid range in Netherlands",
			id:             "456",
			zipcodeFrom:    2000, // Example zip code from the Netherlands
			zipcodeTo:      1000, // Example zip code from the Netherlands
			isoCountry:     "NL", // ISO code for the Netherlands
			expectedResult: model.ZoneRangeModel{},
			expectedError:  fmt.Errorf("invalid zipcode range: zipcodeFrom (%d) cannot be greater than zipcodeTo (%d)", 2000, 1000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult, actualError := createZoneRange(tt.id, tt.zipcodeFrom, tt.zipcodeTo, tt.isoCountry)

			// Check that the expected result matches the actual result
			assert.Equal(t, tt.expectedResult, actualResult)

			if tt.expectedError == nil {
				// If no error is expected, check that the actual error is nil
				assert.NoError(t, actualError)
			} else {
				// If an error is expected, check that the actual error matches the expected error message
				assert.EqualError(t, actualError, tt.expectedError.Error())
			}
		})
	}
}
