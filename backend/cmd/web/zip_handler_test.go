package main

import (
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
}
