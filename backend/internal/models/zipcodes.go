package models

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ZipcodeDatabase = "zipcodes"
	CoordCollection = "coordinates"
)

// ZipCodeDBWrapper is an interface that defines the methods that can be used to interact with the zip code database
// this interface is used to mock the database in the tests
type ZipCodeDBWrapper interface {
	GetZipCodes(reqZipCodeFrom int, reqZipCodeTo int) ([]ZipCode, error)
}

// implements the ZipCodeDBWrapper interface
type ZipCodeDBModel struct {
	DB *mongo.Database
}

type ZipCode struct {
	Code        string      `json:"code" bson:"code"`
	Coordinates interface{} `json:"zone_coordinates" bson:"zone_coordinates"`
	Type        string      `json:"type" bson:"type, omitempty"`
}

// returns a slice of ZipCode structs that are within the range of the zip codes provided
func (z *ZipCodeDBModel) GetZipCodes(reqZipCodeFrom int, reqZipCodeTo int) ([]ZipCode, error) {
	// get the coordinates collection
	coll := z.DB.Collection(CoordCollection)
	// convert the zip codes to strings because the zip codes are stored as strings in the db
	zipFrom := strconv.Itoa(reqZipCodeFrom)
	zipTo := strconv.Itoa(reqZipCodeTo)

	queryFilter := bson.M{"code": bson.M{"$lte": zipTo, "$gte": zipFrom}}
	cur, err := coll.Find(context.TODO(), queryFilter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.TODO())
	// @see https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/cursor/#std-label-golang-cursor
	// TODO: might want to initialize slice with a zise of cur.BatchSize
	// TODO: too many results might cause memory issues... need to investigate -> pagination?
	var results []ZipCode
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
