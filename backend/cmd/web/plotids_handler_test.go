package main

import (
	"bytes"
	"bzone/backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
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

func MockPlotData() bson.D {
	return bson.D{
		{"plot_id", "0"},
		{"plot_name", "Myplot"},

		{"plot_zones", bson.A{
			bson.D{
				{"zone_id", "2"},
				{"zone_ranges", bson.A{
					bson.D{
						{"zipcode_from", 1},
						{"zipcode_to", 4},
						{"iso_country", "NLD"},
						{"zone_range_id", 123},
					},

					bson.D{
						{"zipcode_from", 5},
						{"zipcode_to", 10},
						{"iso_country", "NLD"},
						{"zone_range_id", 2},
					},
				},
				},
				{"zone_fuel_cost", 0},
				{"zone_driving_time", 0},
			},
		}},
		{"plot_created_at", "0001-01-01T00:00:00Z"},
		{"plot_saved_at", "0001-01-01T00:00:00Z"},
	}
}

func MockUsersData() bson.D {
	return bson.D{
		{"user_id", 1},

		{"user_plots", bson.A{
			bson.D{
				{"user_plot_id", "0"},
				{"user_plot_name", "myPlot"},
			},
			bson.D{
				{"user_plot_id", "1"},
				{"user_plot_name", "AccuratePlot"},
			},
		},
		},

		{"uuid", "user2432"},
	}
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
			MockPlotData())

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

	assert.Equal(t, 200, responseValidPlotId.StatusCode) // We expect return code 200

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
			MockUsersData())

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

	// Create your subtest run instance
	mt.Run("MissingUserId", func(mt *mtest.T) {

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for invalid user ID
		usertestCase3(t, app)

	})

	// Create your subtest run instance
	mt.Run("MissingToken", func(mt *mtest.T) {

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing token
		usertestCase4(t, app)

	})

}
func CreateAccessTokenString(userId string) (string, error) {
	// Define the token claims, including the user ID
	claims := jwt.MapClaims{
		"user_id": userId,
	}

	// Define the token signing method and secret key
	signingMethod := jwt.SigningMethodHS256
	secret := []byte("my_secret_key")

	// Create the token
	token := jwt.NewWithClaims(signingMethod, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	// Return the Bearer string access token
	return fmt.Sprintf("Bearer %s", tokenString), nil
}

// usertestCase1
//
//	@Description: test case for valid user and authentication
//	@param t
//	@param a
func usertestCase1(t *testing.T, a *application) {

	// Test handler, valid user ID case
	reqValidUserId := httptest.NewRequest("GET", "/", nil)
	token, _ := CreateAccessTokenString("1")
	reqValidUserId.Header.Add("Authorization", token)
	wValidUserId := httptest.NewRecorder()

	a.getUserPlotIDs(wValidUserId, reqValidUserId)

	responseMissingUserId := wValidUserId.Result()
	bodyValidUserIdJSONString, _ := io.ReadAll(responseMissingUserId.Body)

	// Decode response into object
	var mockDBUser models.UserModel
	json.Unmarshal(bodyValidUserIdJSONString, &mockDBUser)

	assert.Equal(t, 200, responseMissingUserId.StatusCode) // We expect return code 200

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

	// Test handler, valid user ID case
	reqInvalidUserId := httptest.NewRequest("GET", "/", nil)
	token, _ := CreateAccessTokenString("2432")
	reqInvalidUserId.Header.Add("Authorization", token)
	wInalidUserId := httptest.NewRecorder()

	a.getUserPlotIDs(wInalidUserId, reqInvalidUserId)

	responseMissingUserId := wInalidUserId.Result()
	bodyValidUserIdJSONString, _ := io.ReadAll(responseMissingUserId.Body)

	// Decode response into object
	var mockDBUser models.UserModel
	json.Unmarshal(bodyValidUserIdJSONString, &mockDBUser)

	assert.Equal(t, 404, responseMissingUserId.StatusCode) // We expect return code 404

}

// usertestCase3
//
//	@Description: test for missing user id in the token
//	@param t
//	@param a
func usertestCase3(t *testing.T, a *application) {

	// Test handler, missing user ID case
	reqValidUserId := httptest.NewRequest("GET", "/", nil)
	reqValidUserId.Header.Add("Authorization", "Bearer ")
	wValidUserId := httptest.NewRecorder()
	a.getUserPlotIDs(wValidUserId, reqValidUserId)
	responseMissingUserId := wValidUserId.Result()
	assert.Equal(t, 400, responseMissingUserId.StatusCode) // We expect return code 400

}

func usertestCase4(t *testing.T, a *application) {

	// Test handler, missing token
	reqMissingToken := httptest.NewRequest("GET", "/", nil)
	//no authorization token sent
	reqMissingToken.Header.Add("Authorization", "")
	wMissingToken := httptest.NewRecorder()
	a.getUserPlotIDs(wMissingToken, reqMissingToken)
	responseMissingToken := wMissingToken.Result()

	assert.Equal(t, 400, responseMissingToken.StatusCode) // We expect return code 400
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
		//Test case for missing user ID
		userPlotTestCase2(t, app)
		//Test case for missing plot
		userPlotTestCase3(t, app)

	})
	// Create your subtest run instance
	mt.Run("MissingUserIDSavedPlot", func(mt *mtest.T) {

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing user ID
		userPlotTestCase2(t, app)

	})
	// Create your subtest run instance
	mt.Run("MissingPlot", func(mt *mtest.T) {

		//set up the app for the tests
		app := SetUpMockApp(mt)

		//Test case for missing plot
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
	reqSavePlot := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	token, _ := CreateAccessTokenString("1")

	reqSavePlot.Header.Add("Authorization", token)
	wSavePlot := httptest.NewRecorder()
	a.SavePlot(wSavePlot, reqSavePlot)
	responseMissingToken := wSavePlot.Result()

	assert.Equal(t, http.StatusCreated, responseMissingToken.StatusCode) // We expect return code 400
}

// userPlotTestCase2
//
//	@Description: tests saving plot with missing user id
//	@param t
//	@param a
func userPlotTestCase2(t *testing.T, a *application) {

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
	reqSavePlot := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	token, _ := CreateAccessTokenString("")

	reqSavePlot.Header.Add("Authorization", token)
	wSavePlot := httptest.NewRecorder()
	a.SavePlot(wSavePlot, reqSavePlot)
	responseMissingToken := wSavePlot.Result()

	assert.Equal(t, 500, responseMissingToken.StatusCode) // We expect return code 400
}

// userPlotTestCase3
//
//	@Description: sending empty plot model
//	@param t
//	@param a
func userPlotTestCase3(t *testing.T, a *application) {

	var plotModel models.PlotModel

	jsonValue, _ := json.Marshal(plotModel)
	reqSavePlot := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	token, _ := CreateAccessTokenString("")

	reqSavePlot.Header.Add("Authorization", token)
	wSavePlot := httptest.NewRecorder()
	a.SavePlot(wSavePlot, reqSavePlot)
	responseMissingToken := wSavePlot.Result()

	assert.Equal(t, 500, responseMissingToken.StatusCode) // We expect return code 400
}
