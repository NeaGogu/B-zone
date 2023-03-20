package models

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPlotIDs
//
//	@Description: returns the plot IDs and their names that are assigned to the specific user ID
//	@receiver z
//	@param userId
//	@return []PlotIDName
//	@return error
func (z *BzoneDBModel) GetPlotIDs(userId int) ([]PlotIDNamePair, error) {
	// get the users collection
	coll := z.DB.Collection(UserCollection)

	// query for the right user id
	queryFilter := bson.M{"user_id": bson.M{"$eq": userId}}
	var user UserModel

	err := coll.FindOne(context.TODO(), queryFilter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.PlotIdNames, nil
}

func (b *BzoneDBModel) GetPlotById(plotId string) (*PlotModel, error) {

	// get the plots collection
	plotColl := b.DB.Collection(PlotCollection)

	// query by plot_id
	queryFilter := bson.M{"plot_id": bson.M{"$eq": plotId}}

	var plot PlotModel
	err := plotColl.FindOne(context.TODO(), queryFilter).Decode(&plot)
	if err != nil {
		// if the query does not find any result, then FindOne will return a
		// mongo.ErrNoDocuments error. We can check for this error and return
		// our own ErrDocumentNotFound error instead so that it can be handled
		// differently
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrDocumentNotFound
		}

		// otherwise, return the error
		return nil, err
	}

	return &plot, nil
}

func (b *BzoneDBModel) SavePlot(userId int, plot *PlotModel) error {

	// get the users collection and the plots collection
	// userColl := b.DB.Collection(UserCollection)
	plotColl := b.DB.Collection(PlotCollection)

	fmt.Printf("Saving plot: %v", plot)
	// first, add the plot to the plots collection
	res, err := plotColl.InsertOne(context.TODO(), plot)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted plot with ID: %v\n", res.InsertedID)

	// add the plot id to the user's plot ids
	err = b.AddPlotToUserOrCreate(plot, userId)
	if err != nil {
		return err
	}

	return nil
}

// DeletePlots deletes the plots that are assigned to the user based on their origin
// the origin can be "bumba", "algo" or "user"
// returns the number of deleted plots and an error
func (b *BzoneDBModel) DeletePlotsByOrigin(origin string, userId int) (int, error) {

	// get the plots collection
	plotColl := b.DB.Collection(PlotCollection)

	// target all the plots with the origin
	queryFilter := bson.M{
		"origin": bson.M{"$eq": origin},
	}

	// delete the plots
	res, err := plotColl.DeleteMany(context.TODO(), queryFilter)
	if err != nil {
		return 0, err
	}

	// TODO: remove the plot ids from the user's plot ids

	return int(res.DeletedCount), nil
}

// helper function that converts a hex string to an ObjectID that can be used in a MongoDB query
func objectIDFromHex(hex string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objectID, nil
}
