package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetZipCodeCoords(t *testing.T) {
	// Setup mock app
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	app := SetUpMockApp(mt)

	// Test cases
	testCases := []struct {
		name        string
		queryParams string
		expected    int
	}{
		{
			name:        "invalid zip_from",
			queryParams: "zip_from=-1&zip_to=10003",
			expected:    http.StatusBadRequest,
		},
		{
			name:        "invalid zip_to",
			queryParams: "zip_from=10001&zip_to=-1",
			expected:    http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup request and recorder
			req, err := http.NewRequest("GET", "/zip/coordinates?"+tc.queryParams, nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()

			// Call the handler function
			app.getZipCodeCoords(recorder, req)

			// Check the status code
			if recorder.Code != tc.expected {
				t.Errorf("Expected status code %d but got %d", tc.expected, recorder.Code)
			}

		})
	}

	mt.Run("valid zip from and zip to", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.zip",
			mtest.FirstBatch,
			bson.D{{"code", "1"},
				{"zone_coordinates", 1},
				{"type", "type"}})

		findAnother := mtest.CreateCursorResponse(
			1,
			"Bzone.zip",
			mtest.NextBatch,
			bson.D{{"code", "2"},
				{"zone_coordinates", 1},
				{"type", "type"}})

		killCursors := mtest.CreateCursorResponse(0, "Bzone.zip", mtest.NextBatch)

		mt.AddMockResponses(findOne, findAnother, killCursors)
		req := httptest.NewRequest("GET", "/zip/coordinates?zip_from=1&zip_to=3", nil)

		recorder := httptest.NewRecorder()
		app := SetUpMockApp(mt)
		// Call the handler function
		app.getZipCodeCoords(recorder, req)

		// Check the status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
		}

	})

	// Create your subtest run instance
	mt.Run("Error in finding ZipCodes", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})
		req := httptest.NewRequest("GET", "/zip/coordinates?zip_from=1&zip_to=2", nil)

		recorder := httptest.NewRecorder()
		app := SetUpMockApp(mt)
		// Call the handler function
		app.getZipCodeCoords(recorder, req)

		// Check the status code
		if recorder.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, recorder.Code)
		}
	})
}
