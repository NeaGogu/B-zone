package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrDocumentNotFound
		}
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

	// first, add the plot to the plots collection
	_, err := plotColl.InsertOne(context.TODO(), plot)
	if err != nil {
		return err
	}

	// add the plot id to the user's plot ids
	err = b.AddPlotToUserOrCreate(plot, userId)
	if err != nil {
		return err
	}

	return nil
}

// Delete a plot from the plots collection given its id
// returns the number of deleted plots and an error
//
// NOTE: this method also removes the plot id from the user's list of plot ids
func (b *BzoneDBModel) DeletePlotById(plotId string, userId int) (int, error) {

	// get the plots collection
	plotColl := b.DB.Collection(PlotCollection)

	// target the plot with the given plot id
	queryFilter := bson.M{"plot_id": bson.M{"$eq": plotId}}

	res, err := plotColl.DeleteOne(context.TODO(), queryFilter)
	if err != nil {
		return 0, err
	}

	// also remove the plot id from the users's list of plot ids
	_, err = b.deletePlotFromUser(userId, plotId)
	if err != nil {
		return 0, err
	}

	// return the numver of deleted plots, which should be 1
	return int(res.DeletedCount), nil

}

// DeletePlots deletes the plots that are assigned to the user based on their origin
// the origin can be "bumbal", "algo" or "user"
// returns the number of deleted plots and an error
func (b *BzoneDBModel) DeletePlotsByOrigin(origin string, userId int) (int, error) {

	// get the plots collection
	plotColl := b.DB.Collection(PlotCollection)

	// target all the plots with the origin
	queryFilter := bson.M{
		"origin": bson.M{"$eq": origin},
	}

	// NOTE: might want to do all this in a transaction
	// such that if something fails, then the whole thing fails and gets rolled back

	// need to first get the plot ids that are going to be deleted
	// there is no option to return the deleted documents in the DeleteMany method...
	cur, err := plotColl.Find(context.TODO(), queryFilter)
	if err != nil {
		return 0, err
	}

	defer cur.Close(context.TODO())

	var plotsToBeDeleted []PlotModel
	if err = cur.All(context.TODO(), &plotsToBeDeleted); err != nil {
		return 0, err
	}

	// delete the plots
	res, err := plotColl.DeleteMany(context.TODO(), queryFilter)
	if err != nil {
		return 0, err
	}

	// remove the plot ids from the user's plot ids
	for _, plot := range plotsToBeDeleted {
		// this returns the new user model, but we don't need it
		_, err := b.deletePlotFromUser(userId, plot.PlotId)
		if err != nil {
			return 0, err
		}
	}

	return int(res.DeletedCount), nil
}

// DeletePlot deletes the plot with the given plot id from the user's list of plots
// this method is not exported because it's meant to be used internally by
// DeletePlotsByOrigin and DeletePlotById
func (b *BzoneDBModel) deletePlotFromUser(userId int, plotId string) (*UserModel, error) {

	// get the users coll
	userColl := b.DB.Collection(UserCollection)

	// only look at the user with the given user id
	filter := bson.D{{Key: "user_id", Value: userId}}
	// create an update document that removes all the matching PlotIDNamePairs from the
	// user's plots list
	update := bson.D{{Key: "$pull", Value: bson.D{{"user_plots", bson.D{{"user_plot_id", plotId}}}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedUser UserModel
	err := userColl.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

//
//// helper function that converts a hex string to an ObjectID that can be used in a MongoDB query
//func objectIDFromHex(hex string) (primitive.ObjectID, error) {
//	objectID, err := primitive.ObjectIDFromHex(hex)
//	if err != nil {
//		return primitive.NilObjectID, err
//	}
//	return objectID, nil
//}
