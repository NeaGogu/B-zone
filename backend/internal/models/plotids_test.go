package models

import (
	"bzone/backend/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestBzoneDBModel_GetPlotById(t *testing.T) {
	// Create mongo test instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Create your subtest run instance
	mt.Run("ValidPlotId", func(mt *mtest.T) {

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

		//Set up the User data that we expect to get
		UserPlots := []PlotIDNamePair{
			{
				PlotId: "0",
				Name:   "myPlot",
			},

			{
				PlotId: "1",
				Name:   "AccuratePlot",
			},
		}

		//Check if the user data is  the same
		assert.ObjectsAreEqualValues(UserPlots, testUser)

	})

}
