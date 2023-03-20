package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *BzoneDBModel) AddPlotToUserOrCreate(plot *PlotModel, userId int) error {

	// get the users collection
	userColl := b.DB.Collection(UserCollection)

	// Check if a document with the specified user_id exists
	filter := bson.M{"user_id": userId}
	count, err := userColl.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}

	plotNamePair := PlotIDNamePair{
		PlotId: plot.PlotId,
		Name:   plot.Name,
	}

	// If a document with the specified user_id exists, append the new value to the PlotIdNames slice
	if count > 0 {
		update := bson.M{"$push": bson.M{"user_plots": plotNamePair}}
		_, err = userColl.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err
		}

		// TODO: return some value idk
		return nil
	}

	// If a document with the specified user_id does not exist, create a new document with the user_id and the new value in the PlotIdNames slice
	user := UserModel{
		UserId:      int64(userId),
		PlotIdNames: []PlotIDNamePair{plotNamePair},
	}

	_, err = userColl.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	// TODO: return some value idk

	return nil
}
