package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// GetbZnePlot
//
//	@Description: returns the plot based on the plotID
//	@receiver z
//	@param plotid
//	@return BZonePlot
//	@return error
func (z *BzoneDBModel) GetPlot(plotId string) (*PlotModel, error) {
	// get the plots collection
	plotsCollection := z.DB.Collection(PlotCollection)

	//set query filter
	queryFilter := bson.D{{Key: "plot_id", Value: plotId}}

	//Store plot in the PlotModel
	var bZonePlot PlotModel

	//query the database with the set plot id
	err := plotsCollection.FindOne(context.TODO(), queryFilter).Decode(&bZonePlot)
	if err != nil {
		return nil, err
	}

	return &bZonePlot, nil
}
