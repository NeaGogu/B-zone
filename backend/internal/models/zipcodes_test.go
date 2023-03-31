package models

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestZipCodeDBModel_GetZipCodesGetZipCodes(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("Find Zip Codes", func(mt *mtest.T) {

		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.FirstBatch,
			bson.D{{"code", "1"},
				{"zone_coordinates", nil},
				{"type", "type"}})

		findAnother := mtest.CreateCursorResponse(
			1,
			"Bzone.users",
			mtest.NextBatch,
			bson.D{{"code", "2"},
				{"zone_coordinates", nil},
				{"type", "type"}})

		killCursors := mtest.CreateCursorResponse(0, "Bzone.users", mtest.NextBatch)

		mt.AddMockResponses(findOne, findAnother, killCursors)
		var mockZipCodeModel = ZipCodeDBModel{DB: mt.DB}

		_, err := mockZipCodeModel.GetZipCodes(1, 2)
		assert.Nil(t, err)
	})

	// Create your subtest run instance
	mt.Run("Error in finding ZipCodes", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})
		var mockZipCodeModel = ZipCodeDBModel{DB: mt.DB}
		_, err := mockZipCodeModel.GetZipCodes(1, 2)
		assert.NotNil(t, err)
	})
}
