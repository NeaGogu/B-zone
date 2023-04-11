package models

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestBzoneDBModel_AddPlotToUserOrCreate(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Successfully Adding User", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			nil)

		successResponse := mtest.CreateSuccessResponse(
			bson.E{Key: "nModified", Value: 1},
			bson.E{Key: "n", Value: 1},
		)
		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		mt.AddMockResponses(findOne, successResponse, killCursors)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := ValidPlotModelsCollections()
		testUser := mockBZoneModel.AddPlotToUserOrCreate(&plotToSave, 1)

		assert.Nil(t, testUser)

	})

	mt.Run("Successfully Updating User Plots", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"count", 2}, {"ok", 1}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := ValidPlotModelsCollections()
		testUser := mockBZoneModel.AddPlotToUserOrCreate(&plotToSave, 1)

		assert.NotNil(t, testUser)

	})

	mt.Run("Unsuccessfully Finding Plots of User", func(mt *mtest.T) {
		// Mock Plot data
		mt.AddMockResponses(bson.D{{"ok", 0}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := ValidPlotModelsCollections()
		err := mockBZoneModel.AddPlotToUserOrCreate(&plotToSave, 1)
		assert.NotNil(t, err)

	})

	mt.Run("Unsuccessfully Inserting User", func(mt *mtest.T) {
		// Mock Plot data

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			nil)

		mt.AddMockResponses(findOne, bson.D{{"ok", 0}})
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := ValidPlotModelsCollections()
		err := mockBZoneModel.AddPlotToUserOrCreate(&plotToSave, 1)
		assert.NotNil(t, err)

	})

}
