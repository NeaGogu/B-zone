package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// helper method to query the database for a zone by id
// pretty much obsolete because GetZonesListById can behave in the same way
// if passed only one id
func (m *BzoneDBModel) GetZoneById(zoneId string) (*ZoneModel, error) {

	// Get the zones collection
	coll := m.DB.Collection(ZoneCollection)

	// Create a filter with the zone id
	queryFilter := bson.M{"zone_id": zoneId}

	// Create a variable to store the decoded zone
	var zone ZoneModel
	err := coll.FindOne(context.TODO(), queryFilter).Decode(&zone)
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

	return &zone, nil

}

// helper method to query the database for a list of zones
// @deprecated
func (m *BzoneDBModel) GetZonesListById(zoneIds ...string) ([]*ZoneModel, error) {

	// Get the zones collection
	coll := m.DB.Collection(ZoneCollection)

	queryFilter := bson.M{"zone_id": bson.M{"$in": zoneIds}}

	cur, err := coll.Find(context.TODO(), queryFilter)
	if err != nil {
		return nil, err
	}

	// don't forget to close the cursor when you're done
	defer cur.Close(context.TODO())

	var result []*ZoneModel
	if err = cur.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, ErrDocumentNotFound
	}

	return result, nil
}
