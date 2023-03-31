package models

import (
	"bzone/backend/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestBzoneDBModel_GetPlotIDs(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("Valid Plot IDs", func(mt *mtest.T) {

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
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		testUser, _ := mockBZoneModel.GetPlotIDs(1)

		//Check if the user data is  the same
		assert.ObjectsAreEqualValues(ValidUserPlots(), testUser)

	})

	// Create your subtest run instance
	mt.Run("Missing Plot IDs", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		testUser, err := mockBZoneModel.GetPlotIDs(1)

		//Check if the user data is  the same
		assert.ObjectsAreEqual(nil, testUser)
		assert.ObjectsAreEqual(ErrDocumentNotFound, err)

	})

}

func TestBzoneDBModel_GetPlotById(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("Valid PlotId", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			test.MockPlotData())

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		mt.AddMockResponses(findOne, killCursors)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		testPlot, _ := mockBZoneModel.GetPlotById("1")

		//Check if the user data is  the same
		assert.ObjectsAreEqualValues(ValidPlotModelsCollections(), testPlot)

	})

	// Create your subtest run instance
	mt.Run("Missing PlotId", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		testUser, err := mockBZoneModel.GetPlotById("1")

		//Check if the user data is  the same
		assert.ObjectsAreEqual(nil, testUser)
		assert.ObjectsAreEqual(ErrDocumentNotFound, err)

	})

}

func TestBzoneDBModel_SavePlot(t *testing.T) {
	// Create your subtest run instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Successfully Saved Plot", func(mt *mtest.T) {

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

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := ValidPlotModelsCollections()
		testSavePlot := mockBZoneModel.SavePlot(1, &plotToSave)
		assert.ObjectsAreEqual(testSavePlot, nil)
	})

	mt.Run("Unsuccessfully Saved Plot", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := PlotModel{}
		err := mockBZoneModel.SavePlot(1, &plotToSave)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestBzoneDBModel_deletePlotFromUser(t *testing.T) {
	// Create your subtest run instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Successfully Deleted Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.deletePlotFromUser(1, "0")
		assert.ObjectsAreEqual(err, nil)
	})

	mt.Run("Unsuccessfully Deleted Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.deletePlotFromUser(1, "0")
		assert.NotNil(t, err)
	})
}

func TestBzoneDBModel_DeletePlotsByOrigin(t *testing.T) {
	// Create your subtest run instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Successfully Deleted Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotsByOrigin("Bumbal", 1)
		assert.ObjectsAreEqual(err, nil)
	})

	mt.Run("Unsuccessfully Deleted Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotsByOrigin("Bumbal", 1)
		assert.NotNil(t, err)
	})
}

func TestBzoneDBModel_DeletePlotById(t *testing.T) {
	// Create your subtest run instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Successfully Deleted Plot By ID", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotById("0", 1)
		assert.ObjectsAreEqual(err, nil)
	})

	mt.Run("Unsuccessfully Deleted Plot By ID", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotById("0", 1)
		assert.NotNil(t, err)
	})
}
