package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ZoneRangeDBModel struct {
	DB *mongo.Database
}

// ZoneRangeModel struct for ZoneRangeModel
type ZoneRangeModel struct {
	// Unique Zone type ID
	ZoneRangeId int64 `json:"zone_range_id,omitempty"`
	// Zipcode range start
	ZipcodeFrom int64 `json:"zipcode_from,omitempty"`
	// Zipcode range end
	ZipcodeTo int64 `json:"zipcode_to,omitempty"`
	// iso country of the zone range
	IsoCountry string `json:"iso_country,omitempty"`
	//array of coordinates?
}
