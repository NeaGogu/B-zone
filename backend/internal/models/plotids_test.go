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
	mt.Run("Missing Document of Plot IDs", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch)

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		//Create Mock Responses
		//The given bson.D will be returned from the mongo to the driver

		mt.AddMockResponses(findOne, killCursors)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.GetPlotIDs(1)

		//Check if the user data is  the same
		assert.NotNil(t, err)

	})

	// Create your subtest run instance
	mt.Run("Missing Plot IDs", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.GetPlotIDs(1)

		//Check if the user data is  the same
		assert.NotNil(t, err)

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
		_, err := mockBZoneModel.GetPlotById("1")

		//Check if the user data is  the same
		assert.Nil(t, err)

	})

	// Create your subtest run instance
	mt.Run("Missing Plot Document for PlotId", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch)

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		mt.AddMockResponses(findOne, killCursors)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.GetPlotById("1")

		//Check if the user data is  the same
		assert.NotNil(t, err)

	})

	// Create your subtest run instance
	mt.Run("Missing PlotId", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.GetPlotById("1")

		assert.NotNil(t, err)

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
		assert.Nil(t, testSavePlot)
	})

	mt.Run("Unsuccessfully Founded Plot", func(mt *mtest.T) {

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

	mt.Run("Unsuccessfully Created Plot", func(mt *mtest.T) {

		successResponse := mtest.CreateSuccessResponse(
			bson.E{Key: "nModified", Value: 1},
			bson.E{Key: "n", Value: 1},
		)
		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)
		mt.AddMockResponses(successResponse, killCursors)

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := PlotModel{}
		err := mockBZoneModel.SavePlot(1, &plotToSave)
		assert.NotNil(t, err)
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

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, test.MockPlotData())
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		deletePlot := bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 2}}
		deleteUserPlot := bson.D{{"ok", 1}, {"value", test.MockPlotData()}}

		mt.AddMockResponses(first, killCursors, deletePlot, deleteUserPlot)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotsByOrigin("Bumbal", 1)
		assert.Nil(t, err)
	})

	mt.Run("Unsuccessfully Deleted Plot From Users", func(mt *mtest.T) {

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, test.MockPlotData())
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, test.MockPlotData())
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)

		deletePlot := bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 2}}
		deleteUserPlot := bson.D{{"ok", 1}, {"value", test.MockPlotData()}}

		mt.AddMockResponses(first, second, killCursors, deletePlot, deleteUserPlot)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotsByOrigin("Bumbal", 1)
		assert.NotNil(t, err)
	})

	mt.Run("Unsuccessful Find of Deleted Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotsByOrigin("Bumbal", 1)
		assert.NotNil(t, err)
	})

	mt.Run("Cursor Error Delete Plot", func(mt *mtest.T) {

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, test.MockPlotData())
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, test.MockPlotData())
		deletePlot := bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 2}}
		deleteUserPlot := bson.D{{"ok", 1}, {"value", test.MockPlotData()}}

		mt.AddMockResponses(first, second, deletePlot, deleteUserPlot)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotsByOrigin("Bumbal", 1)
		assert.NotNil(t, err)
	})

	mt.Run("Unsuccessfully Deleted Plot", func(mt *mtest.T) {

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, test.MockPlotData())
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, test.MockPlotData())
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		deletePlot := bson.D{{"ok", 0}, {"acknowledged", true}, {"n", 2}}

		mt.AddMockResponses(first, second, killCursors, deletePlot)
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
		}, bson.D{
			{"ok", 1},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotById("0", 1)
		assert.Nil(t, err)
	})

	mt.Run("Unsuccessfully Deleted Plot By ID from User Collection", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", test.MockPlotData()},
		})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.DeletePlotById("0", 1)
		assert.NotNil(t, err)
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
