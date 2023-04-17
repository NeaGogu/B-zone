package models

import (
	"bzone/backend/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestBzoneDBModel_GetPlot(t *testing.T) {
	// Create your subtest run instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Successfully Retrieved Plot", func(mt *mtest.T) {
		// Mock Plot data
		findOne := mtest.CreateCursorResponse(
			1,
			"Bzone.plots",
			mtest.FirstBatch,
			test.MockPlotData())

		killCursors := mtest.CreateCursorResponse(0, "Bzone.plots", mtest.NextBatch)

		mt.AddMockResponses(findOne, killCursors)

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.GetPlot("0")
		assert.ObjectsAreEqual(err, nil)
	})

	mt.Run("UnSuccessfully Retrieved Plot", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{"ok", 0}})

		var mockBZoneModel = BzoneDBModel{DB: mt.DB}
		_, err := mockBZoneModel.GetPlot("0")
		assert.NotNil(t, err)
	})
}
