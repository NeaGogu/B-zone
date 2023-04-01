package main

import (
	"bytes"
	"bzone/backend/internal/models"
	"bzone/backend/internal/test"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// SetUpMockApp
//
//	@Description: set up the application for testing purposes
//	@param mt
//	@return *application
func SetUpMockApp(mt *mtest.T) *application {
	// Setup application with mocks for test purposes.
	// Set up the test environment
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// This could be re-used for multiple tests.
	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		zipCodeDbModel: nil,
		bzoneDbModel:   &models.BzoneDBModel{DB: mt.DB},
	}
	return app
}

// TestGetPlotBy
//
//	@Description: tests plotids handler
//	@param t
func TestApplication_GetPlotById(t *testing.T) {

	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("ValidPlotId", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.plots",
			mtest.FirstBatch,
			test.MockPlotData())

		killCursors := mtest.CreateCursorResponse(0, "Bzone.plots", mtest.NextBatch)

		//Create Mock Responses
		//The given bson.D will be returned from the mongo to the driver
		mt.AddMockResponses(findOne, killCursors)
		mt.AddMockResponses(bson.D{{"ok", 0}})
		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for valid plot ID
		testCase1(t, app)

	})

	// Create your subtest for missing plot ID
	mt.Run("MissingPlotId", func(mt *mtest.T) {

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing plot ID
		testCase2(t, app)

	})
	// Create your subtest for missing plot ID
	mt.Run("InvalidPlotId", func(mt *mtest.T) {
		// Mock Plot data with empty document
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.plots",
			mtest.FirstBatch,
			nil)

		killCursors := mtest.CreateCursorResponse(0, "Bzone.plots", mtest.NextBatch)
		mt.AddMockResponses(findOne, killCursors)

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for error in mongo DB
		testCase4(t, app)

	})

	// Create your subtest for error message in mongo
	mt.Run("ErrorInDataBase", func(mt *mtest.T) {
		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.plots",
			mtest.FirstBatch)

		killCursors := mtest.CreateCursorResponse(0, "Bzone.plots", mtest.NextBatch)
		mt.AddMockResponses(findOne, killCursors)

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing plot ID
		testCase3(t, app)

	})
}

// testCase1
//
//	@Description: Valid plot ID
//	@param t
//	@param a
func testCase1(t *testing.T, a *application) {

	// Test handler, Valid plot ID
	reqValidPlotId, _ := http.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	reqValidPlotId = reqValidPlotId.WithContext(context.WithValue(reqValidPlotId.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("plotId", "1")

	wValidPlotId := httptest.NewRecorder()
	a.GetPlotById(wValidPlotId, reqValidPlotId)

	responseValidPlotId := wValidPlotId.Result()
	bodyValidPlotIdJSONString, _ := io.ReadAll(responseValidPlotId.Body)

	assert.Equal(t, http.StatusOK, responseValidPlotId.StatusCode) // We expect return code 200

	// Decode response into object
	var bZonePlot models.PlotModel
	json.Unmarshal(bodyValidPlotIdJSONString, &bZonePlot)

	//Set up the data that we expect to get
	//Set up ZoneRanges
	ZoneRanges := []models.ZoneRangeModel{
		{
			ZoneRangeId: 1,
			ZipcodeFrom: 2,
			ZipcodeTo:   4,
			IsoCountry:  "NLD",
		},

		{
			ZoneRangeId: 2,
			ZipcodeFrom: 5,
			ZipcodeTo:   10,
			IsoCountry:  "NLD",
		},
	}

	//Mock Zones
	var mockZones = []models.ZoneModel{
		{
			Id:              "2",
			ZoneRanges:      ZoneRanges,
			ZoneFuelCost:    0,
			ZoneDrivingTime: 0,
		},
	}

	//Check if the plot name is the same
	assert.Equal(t, "Myplot", bZonePlot.Name)

	//Check if the zones are the same
	assert.ObjectsAreEqualValues(mockZones, bZonePlot.Zones)

}

// testCase2
//
//	@Description: Missing plot ID
//	@param t
//	@param a
func testCase2(t *testing.T, a *application) {

	// Test handler, Missing plot ID case
	reqMissingPlotId := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	reqMissingPlotId = reqMissingPlotId.WithContext(context.WithValue(reqMissingPlotId.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("plotId", "")
	wMissingPlotId := httptest.NewRecorder()
	a.GetPlotById(wMissingPlotId, reqMissingPlotId)

	responseMissingPlotId := wMissingPlotId.Result()
	body, _ := io.ReadAll(responseMissingPlotId.Body)

	assert.Equal(t, 400, responseMissingPlotId.StatusCode) // We expect return code 400
	assert.Equal(t, "Plot ID is required\n", string(body)) // We expect "Invalid plot id in query"
}

// testCase3
//
//	@Description: Plot is not found in the DB
//	@param t
//	@param a
func testCase3(t *testing.T, a *application) {

	// Test handler, Invalid plot ID case
	reqInvalidPlotId := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	reqInvalidPlotId = reqInvalidPlotId.WithContext(context.WithValue(reqInvalidPlotId.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("plotId", "invalidID")
	wInvalidPlotId := httptest.NewRecorder()
	a.GetPlotById(wInvalidPlotId, reqInvalidPlotId)

	responseMissingPlotId := wInvalidPlotId.Result()
	body, _ := io.ReadAll(responseMissingPlotId.Body)

	assert.Equal(t, 404, responseMissingPlotId.StatusCode) // We expect return code 404
	assert.Equal(t, "Not Found\n", string(body))
}

// testCase4
//
//	@Description: Error occured in DB
//	@param t
//	@param a
func testCase4(t *testing.T, a *application) {

	// Test handler, Invalid plot ID case
	reqInvalidPlotId := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	reqInvalidPlotId = reqInvalidPlotId.WithContext(context.WithValue(reqInvalidPlotId.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("plotId", "errorID")
	wInvalidPlotId := httptest.NewRecorder()
	a.GetPlotById(wInvalidPlotId, reqInvalidPlotId)

	responseMissingPlotId := wInvalidPlotId.Result()
	body, _ := io.ReadAll(responseMissingPlotId.Body)

	assert.Equal(t, 500, responseMissingPlotId.StatusCode) // We expect return code 500
	assert.Equal(t, "Internal Server Error\n", string(body))
}

func TestApplication_getUserPlotIDs(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("ValidUserPlotIDs", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			test.MockUsersData())

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		//Create Mock Responses
		//The given bson.D will be returned from the mongo to the driver
		mt.AddMockResponses(findOne, killCursors)
		mt.AddMockResponses(bson.D{{"ok", 0}})
		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for valid user ID
		usertestCase1(t, app)

	})
	// Create your subtest run instance
	mt.Run("InvalidUserPlotIDs", func(mt *mtest.T) {

		// Mock User data with empty document
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch)

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)
		mt.AddMockResponses(findOne, killCursors)

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for invalid user ID
		usertestCase2(t, app)

	})

}

// usertestCase1
//
//	@Description: test case for valid user and authentication
//	@param t
//	@param a
func usertestCase1(t *testing.T, a *application) {

	// Test handler, valid user ID case
	req := httptest.NewRequest("GET", "/", nil)
	ctx := context.WithValue(req.Context(), ContextUserKey, 1)
	req = req.WithContext(ctx)
	wValidUserId := httptest.NewRecorder()
	a.getUserPlotIDs(wValidUserId, req)

	responseMissingUserId := wValidUserId.Result()
	bodyValidUserIdJSONString, _ := io.ReadAll(responseMissingUserId.Body)

	// Decode response into object
	var mockDBUser models.UserModel
	json.Unmarshal(bodyValidUserIdJSONString, &mockDBUser)

	assert.Equal(t, http.StatusOK, responseMissingUserId.StatusCode) // We expect return code 200

	//Set up the User data that we expect to get
	UserPlots := []models.PlotIDNamePair{
		{
			PlotId: "0",
			Name:   "myPlot",
		},

		{
			PlotId: "1",
			Name:   "AccuratePlot",
		},
	}

	//Mock Zones
	var UserData = []models.UserModel{
		{
			UserId:      1,
			PlotIdNames: UserPlots,
			Uuid:        "user2432",
		},
	}

	//Check if the user data is  the same
	assert.ObjectsAreEqualValues(UserData, mockDBUser)
}

// usertestCase2
//
//	@Description: tests for invalid user id - do document foung
//	@param t
//	@param a
func usertestCase2(t *testing.T, a *application) {

	// Test handler, invalid user ID case
	req := httptest.NewRequest("GET", "/", nil)
	ctx := context.WithValue(req.Context(), ContextUserKey, 22)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	a.getUserPlotIDs(w, req)

	responseMissingUserId := w.Result()
	bodyValidUserIdJSONString, _ := io.ReadAll(responseMissingUserId.Body)

	// Decode response into object
	var mockDBUser models.UserModel
	json.Unmarshal(bodyValidUserIdJSONString, &mockDBUser)

	assert.Equal(t, http.StatusNotFound, responseMissingUserId.StatusCode) // We expect return code 404

}

func TestApplication_SavePlot(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("SuccessfulSavedPlot", func(mt *mtest.T) {

		successResponse := mtest.CreateSuccessResponse(
			bson.E{Key: "nModified", Value: 1},
			bson.E{Key: "n", Value: 1},
		)
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			nil)

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)
		mt.AddMockResponses(successResponse, findOne, killCursors)
		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for valid user ID
		userPlotTestCase1(t, app)

	})

	// Create your subtest run instance
	mt.Run("MissingPlot", func(mt *mtest.T) {

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing plot
		userPlotTestCase2(t, app)

		//Test case missing jsonBody in request
		userPlotTestCase3(t, app)

	})

}

// userPlotTestCase1
//
//	@Description: Saves plot
//	@param t
//	@param a
func userPlotTestCase1(t *testing.T, a *application) {

	plotModel := models.PlotModel{
		PlotId:        "1",
		Name:          "SavedPlot",
		ZoneIds:       nil,
		Zones:         nil,
		PlotCreatedAt: time.Time{},
		PlotSavedAt:   time.Time{},
		Origin:        "Bumbal",
	}

	jsonValue, _ := json.Marshal(plotModel)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	ctx := context.WithValue(req.Context(), ContextUserKey, 1)
	req = req.WithContext(ctx)

	wSavePlot := httptest.NewRecorder()
	a.SavePlot(wSavePlot, req)
	responsePlotCreated := wSavePlot.Result()

	assert.Equal(t, http.StatusCreated, responsePlotCreated.StatusCode)
}

// userPlotTestCase3
//
//	@Description: sending empty plot model
//	@param t
//	@param a
func userPlotTestCase2(t *testing.T, a *application) {

	var plotModel models.PlotModel

	jsonValue, _ := json.Marshal(plotModel)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	ctx := context.WithValue(req.Context(), ContextUserKey, 22)
	req = req.WithContext(ctx)
	wSavePlot := httptest.NewRecorder()
	a.SavePlot(wSavePlot, req)
	responseEmptyPlot := wSavePlot.Result()

	assert.Equal(t, http.StatusInternalServerError, responseEmptyPlot.StatusCode)
}

// userPlotTestCase3
//
//	@Description: missing plot
//	@param t
//	@param a
func userPlotTestCase3(t *testing.T, a *application) {
	req := httptest.NewRequest("POST", "/", nil)
	ctx := context.WithValue(req.Context(), ContextUserKey, 22)
	req = req.WithContext(ctx)
	wSavePlot := httptest.NewRecorder()
	a.SavePlot(wSavePlot, req)
	responseEmptyPlot := wSavePlot.Result()

	assert.Equal(t, http.StatusBadRequest, responseEmptyPlot.StatusCode)
}

func TestApplication_DeletePlotById(t *testing.T) {
	// Create your subtest run instance

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Server Error when deleting Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for invalid token
		deletePlotTestCase0(t, app)

	})
	mt.Run("Bad Request config for deleting plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing plot id
		deletePlotTestCase1(t, app)

	})

}

func deletePlotTestCase0(t *testing.T, a *application) {

	req, _ := http.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	ctx := context.WithValue(req.Context(), ContextUserKey, 1)
	req = req.WithContext(ctx)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("plotId", "1")
	wSavePlot := httptest.NewRecorder()
	a.DeletePlotById(wSavePlot, req)
	responseFailed := wSavePlot.Result()

	assert.Equal(t, http.StatusInternalServerError, responseFailed.StatusCode)
}

func deletePlotTestCase1(t *testing.T, a *application) {

	req, _ := http.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	ctx := context.WithValue(req.Context(), ContextUserKey, 1)
	req = req.WithContext(ctx)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("plotId", "")
	wSavePlot := httptest.NewRecorder()
	a.DeletePlotById(wSavePlot, req)
	responseMissingID := wSavePlot.Result()

	assert.Equal(t, http.StatusBadRequest, responseMissingID.StatusCode)
}
