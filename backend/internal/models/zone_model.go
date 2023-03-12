package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ZoneDBModel struct {
	DB *mongo.Database
}

// ZoneModel struct for ZoneModel
type ZoneModel struct {
	// Unique Zone ID
	Id int64 `json:"zone_id,omitempty"`
	//Zones Ranges
	ZoneRanges []ZoneRangeModel `json:"zone_ranges,omitempty"`
	//Total fuel cost
	ZoneFuelCost float64 `json:"zone_fuel_cost,omitempty"`
	//Total driving time
	ZoneDrivingTime int64 `json:"zone_driving_time,omitempty"`
}
