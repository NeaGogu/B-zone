package main

import (
	"bzone/backend/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mockZones []interface{} = []interface{} {
	models.ZoneModel{
		Id: uuid.New().String(),
		ZoneFuelCost: 1.5,
		ZoneDrivingTime: 10,
		ZoneRanges: []models.ZoneRangeModel{

			models.ZoneRangeModel{
				ZipcodeFrom: 0,
				ZipcodeTo: 10,
			},
			models.ZoneRangeModel{
				ZipcodeFrom: 11,
				ZipcodeTo: 20,
			},
			models.ZoneRangeModel{
				ZipcodeFrom: 21,
				ZipcodeTo: 30,
			},
		},
	},
	models.ZoneModel{
		Id: uuid.New().String(),
		ZoneFuelCost: 2.5,
		ZoneDrivingTime: 20,
		ZoneRanges: []models.ZoneRangeModel{

			models.ZoneRangeModel{
				ZipcodeFrom: 31,
				ZipcodeTo: 40,
			},
			models.ZoneRangeModel{
				ZipcodeFrom: 41,
				ZipcodeTo: 50,
			},
		},
	},
}

// BE CAREFUL WITH THIS FUNCTION, IT WILL DELETE ALL THE DATA IN THE DATABASE
// ONLY USE IT TO POPULATE THE DATABASE
func populateZonesDummy(uri string) {
	
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Disconnect the client once the function returns
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	zonesColl := client.Database(models.BzoneDatabase).Collection(models.ZoneCollection)

	res, err := zonesColl.InsertMany(context.TODO(), mockZones)
	if err != nil {
		panic(err)
	}


	fmt.Printf("Inserted %v documents into the collection!\n", len(res.InsertedIDs))
}
