package models

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const BzoneDatabase = "Bzone"

// collection names for the bzone database
const (
	UserCollection = "users"
	PlotCollection = "plots"
	ZoneCollection = "zones"
)

// BzoneDBModel is a wrapper for the mongo database that refers to the bzone database
type BzoneDBModel struct {
	DB *mongo.Database
}

// PlotModel struct for ZonesPLotModel
type PlotModel struct {
	// Unique Plot ID
	PlotId int64 `json:"plot_id,omitempty" bson:"plot_id,omitempty"`
	// Plot Name
	Name string `json:"plot_name,omitempty" bson:"plot_name,omitempty"`
	//Zones in the plot
	ZoneIds []int64 `json:"plot_zone_ids,omitempty" bson:"plot_zone_ids,omitempty"`
	//created_at date time
	PlotCreatedAt time.Time `json:"plot_created_at,omitempty" bson:"plot_created_at,omitempty"`
	//saved_at date time
	PlotSavedAt time.Time `json:"plot_saved_at,omitempty" bson:"plot_saved_at,omitempty"`
}

// ZoneModel struct for ZoneModel
type ZoneModel struct {
	// Unique Zone ID
	Id int64 `json:"zone_id,omitempty" bson:"zone_id,omitempty"`
	//Zones Ranges
	ZoneRanges []ZoneRangeModel `json:"zone_ranges,omitempty" bson:"zone_ranges,omitempty"`
	//Total fuel cost
	ZoneFuelCost float64 `json:"zone_fuel_cost,omitempty" bson:"zone_fuel_cost,omitempty"`
	//Total driving time
	ZoneDrivingTime int64 `json:"zone_driving_time,omitempty" bson:"zone_driving_time,omitempty"`
}

// PlotIDNamePair struct containing the ID of a plot and its name
type PlotIDNamePair struct {
	// Unique Plot ID
	PlotId int64 `json:"user_plot_id,omitempty" bson:"user_plot_id,omitempty"`
	// Plot Name
	Name string `json:"user_plot_name,omitempty" bson:"user_plot_name,omitempty"`
}

// UserModel struct for UsersModel
type UserModel struct {
	//B-Zone User ID
	UserId int64 `json:"user_id,omitempty" bson:"user_id,omitempty"`
	// unique per user
	Uuid string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	// plots per user
	PlotIdNames []PlotIDNamePair `json:"user_id_name,omitempty" bson:"user_id_name,omitempty"`
}

// ZoneRangeModel struct for ZoneRangeModel
type ZoneRangeModel struct {
	// Unique Zone type ID
	ZoneRangeId int64 `json:"zone_range_id,omitempty" bson:"zone_range_id,omitempty"`
	// Zipcode range start
	ZipcodeFrom int64 `json:"zipcode_from,omitempty" bson:"zipcode_from,omitempty"`
	// Zipcode range end
	ZipcodeTo int64 `json:"zipcode_to,omitempty" bson:"zipcode_to,omitempty"`
	// iso country of the zone range
	IsoCountry string `json:"iso_country,omitempty" bson:"iso_country,omitempty"`
	//array of coordinates?
}


