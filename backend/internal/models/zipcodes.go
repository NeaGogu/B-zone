package models

import "go.mongodb.org/mongo-driver/mongo"

type ZipCodeDBModel struct {
	DB *mongo.Database
}

type ZipCode struct {
	Code string `json:"code" bson:"code"`
	// PolygonCoordinates [][][]float64 `json:"zone_coordinates" bson:"zone_coordinates, omitempty"`
	// MultiPolyCoordinates [][][][]float64 `json:"zone_coordinates" bson:"zone_coordinates, omitempty"`
	Coordinates interface{} `json:"zone_coordinates" bson:"zone_coordinates"`
	Type string `json:"type" bson:"type, omitempty"`
}
