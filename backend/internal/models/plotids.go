package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

//type PlotIDName struct {
//	PlotId int64  `json:"user_plot_id" bson:"user_plot_id"`
//	Name   string `json:"user_plot_name" bson:"user_plot_name"`
//}

type PlotIDName struct {
	IDNamePairs []PlotIDNamePair `json:"user_id_name" bson:"user_id_name"`
}

// GetPlotIDs
//
//	@Description: returns the plot IDs and their names that are assigned to the specific user ID
//	@receiver z
//	@param userId
//	@return []PlotIDName
//	@return error
func (z *BzoneDBModel) GetPlotIDs(userId int) ([]PlotIDName, error) {
	// get the users collection
	coll := z.DB.Collection(UserCollection)

	// query for the right user id
	queryFilter := bson.M{"user_id": bson.M{"$eq": userId}}
	cur, err := coll.Find(context.TODO(), queryFilter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.TODO())

	// put the plot ids and their name in the results variable
	var results []PlotIDName
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
