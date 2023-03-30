package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ValidUserPlots() []PlotIDNamePair {
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
	return UserPlots
}

func ValidUserCollectionData() []UserModel {
	var userData = []UserModel{
		{
			UserId:      1,
			PlotIdNames: ValidUserPlots(),
			Uuid:        "testUser",
		},
	}
	return userData
}

func SucessResponseSavePlot() (bson.E, bson.E) {
	return bson.E{Key: "nModified", Value: 1}, bson.E{Key: "n", Value: 1}
}

func ValidPlotModelsCollections() PlotModel {
	return PlotModel{
		PlotId:        "1",
		Name:          "SavedPlot",
		ZoneIds:       nil,
		Zones:         nil,
		PlotCreatedAt: time.Time{},
		PlotSavedAt:   time.Time{},
		Origin:        "Bumbal",
	}

}
