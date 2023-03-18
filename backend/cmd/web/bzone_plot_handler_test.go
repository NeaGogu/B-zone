package main

import (
	"bzone/backend/internal/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func SetUpMockApp(mt *mtest.T) *application {
	// Setup application with mocks for test purposes.
	// Set up the test environment
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// This could be re-used for multiple tests.
	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		zipCodeDbModel: &models.ZipCodeDBModel{DB: mt.DB},
		bzoneDbModel:   &models.BzoneDBModel{DB: mt.DB},
	}
	return app
}

func TestBzonePlotHandler(t *testing.T) {

	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("PlotsTest", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.plots",
			mtest.FirstBatch,
			bson.D{{"_id", primitive.NewObjectID()},
				{"plot_id", "0"},
				{"plot_name", "plot0"},
				{"plot_zone_ids", bson.A{0}},
				{}})

		killCursors := mtest.CreateCursorResponse(0, "Bzone.plots", mtest.NextBatch)

		//Create Mock Responses
		//The given bson.D will be returned from the mongo to the driver
		mt.AddMockResponses(findOne, killCursors)

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing plot ID
		testCase1(t, app)

		//Test case for correct plot ID
		testCase2(t, app)

		//Test case for invalid plot ID
		testCase3(t, app)

	})

}

// Missing plot ID
func testCase1(t *testing.T, a *application) {

	// Test handler, Invalid plot ID case
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	a.getBZonePlot(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)                       // We expect return code 400
	assert.Equal(t, "Invalid plot id in query\n", string(body)) // We expect "Invalid plot id in query"
}

// Valid plot ID
func testCase2(t *testing.T, a *application) {

	// Test handler, Valid plot ID
	reqValidPlotId := httptest.NewRequest("GET", "/test?plot_id=0", nil)
	wValidPlotId := httptest.NewRecorder()
	a.getBZonePlot(wValidPlotId, reqValidPlotId)

	responseValidPlotId := wValidPlotId.Result()
	bodyValidPlotIdJSONString, _ := io.ReadAll(responseValidPlotId.Body)

	assert.Equal(t, 200, responseValidPlotId.StatusCode) // We expect return code 200

	// Print JSON response
	t.Log(string(bodyValidPlotIdJSONString))

	// Decode response into object
	var bZonePlot models.PlotModel
	json.Unmarshal(bodyValidPlotIdJSONString, &bZonePlot)

	// Check object correctness based on our previous mocks
	assert.Equal(t, "plot0", bZonePlot.Name)
	assert.Equal(t, []int64{0}, bZonePlot.ZoneIds)
}

// Plot is not found in the DB
func testCase3(t *testing.T, a *application) {

	// Test handler, Valid plot ID
	req := httptest.NewRequest("GET", "/test?plot_id=3", nil)
	w := httptest.NewRecorder()
	a.getBZonePlot(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 404, resp.StatusCode)                  // We expect return code 404
	assert.Equal(t, "Plot does not exist\n", string(body)) // We expect "Plot does not exist"
}
