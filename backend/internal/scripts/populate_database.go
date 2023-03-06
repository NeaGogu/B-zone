package scripts

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	geojson "github.com/paulmach/go.geojson"
)

const uri = "mongodb://localhost:27017"


type ZipCode struct {
	Code string `json:"code" bson:"code"`
	// PolygonCoordinates [][][]float64 `json:"zone_coordinates" bson:"zone_coordinates, omitempty"`
	// MultiPolyCoordinates [][][][]float64 `json:"zone_coordinates" bson:"zone_coordinates, omitempty"`
	Coordinates interface{} `json:"zone_coordinates" bson:"zone_coordinates"`
	Type string `json:"type" bson:"type, omitempty"`
}


// BE CAREFUL WITH THIS FUNCTION, IT WILL DELETE ALL THE DATA IN THE DATABASE
// ONLY USE IT TO POPULATE THE DATABASE
func populateDatabase() {
	
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

	zipCoordCollection := client.Database("zipcodes").Collection("coordinates")
	
	// First clean the collection 
	fmt.Println("[!!] Deleting all documents in the collection")
	zipCoordCollection.DeleteMany(context.TODO(), bson.D{})	

	// result, err := zipCoordCollection.InsertOne(context.TODO(), testDoc)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	
	// Read the geojson file

	// TODO provide this as argument
	dataFilePath := "/home/gogu/Projects/SEP/test_geoson/data_inverted.json"

	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("opened file: %s\n", dataFilePath)

	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// loop through features and add them to database

	fmt.Print("----------Inserting zipcodes into coordinates collection-------\n")
	var zipDocument ZipCode

	for _, f := range fc.Features {
		featureGeometry := f.Geometry
		
		zipDocument.Code = f.Properties["pc4_code"].(string)
		zipDocument.Coordinates = nil
		zipDocument.Type = string(featureGeometry.Type)

		// invert coordinates
		switch {
		case featureGeometry.IsPolygon():
			fmt.Printf("Inserting zipcode %s of type %v\n", zipDocument.Code, featureGeometry.Type)
			zipDocument.Coordinates = featureGeometry.Polygon
			
			result, err := zipCoordCollection.InsertOne(context.TODO(), zipDocument)

			if err != nil {
				panic(err)
			}

			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)


		case featureGeometry.IsMultiPolygon():
			fmt.Printf("Inserting zipcode %s of type %v\n", zipDocument.Code, featureGeometry.Type)
			zipDocument.Coordinates = featureGeometry.MultiPolygon

			result, err := zipCoordCollection.InsertOne(context.TODO(), zipDocument)

			if err != nil {
				panic(err)
			}

			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

		default:
			fmt.Fprintf(os.Stderr, "%v is neither polygon nor multipolygon, will not add\n", zipDocument.Code) 

		}

	}

	fmt.Print("----------Done inserting zipcodes into coordinates collection-------\n")
	
}

