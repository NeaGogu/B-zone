package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const(
	CoordCollection = "coordinates"
)

type ZipCodeDBModel struct {
	DB *mongo.Database
}

type ZipCode struct {
	Code string `json:"code" bson:"code"`
	Coordinates interface{} `json:"zone_coordinates" bson:"zone_coordinates"`
	Type string `json:"type" bson:"type, omitempty"`
}


	// queryFilter := bson.D{{Key: "code", Value: reqZipCode}}
	//
	// // TODO add this collection as a constant
	// coll := app.zipCodeDbModel.DB.Collection("coordinates")
	//
	// var zipCode models.ZipCode
	// err := coll.FindOne(context.TODO(), queryFilter).Decode(&zipCode)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

func (z *ZipCodeDBModel) GetZipCodes(reqZipCodeFrom string, reqZipCodeTo string) ([]ZipCode, error) {
	// get the coordinates collection
	coll := z.DB.Collection(CoordCollection)


	queryFilter := bson.M { "code" : bson.M {"$lte": reqZipCodeTo, "$gte": reqZipCodeFrom}}
	cur, err:= coll.Find(context.TODO(), queryFilter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.TODO())
	// @see https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/cursor/#std-label-golang-cursor
	// TODO might want to initialize slice with a zise of cur.BatchSize
	// TODO too many results might cause memory issues... need to investigate -> pagination?
	var results []ZipCode
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
