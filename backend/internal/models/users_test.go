package models

import (
	"bzone/backend/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestBzoneDBModel_AddPlotToUserOrCreate(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("Update User", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			test.MockUsersData())

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		mt.AddMockResponses(findOne, killCursors)
		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		plotToSave := ValidPlotModelsCollections()
		testUser := mockBZoneModel.AddPlotToUserOrCreate(&plotToSave, 1)
		assert.ObjectsAreEqual(testUser, nil)

	})
	mt.Run("Add User", func(mt *mtest.T) {

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

		assert.ObjectsAreEqual(testUser, nil)

	})

	mt.Run("Unsucessfully Adding User", func(mt *mtest.T) {

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

		assert.ObjectsAreEqual(testUser, nil)

	})
}
