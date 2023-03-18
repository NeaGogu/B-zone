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

// origin of data: bumbal, algorithm, bzone
const (
	OriginBumbal = "bumbal"
	OriginBzone  = "bzone"
	OriginAlgo   = "algo"
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
	// @deprecated use Zones instead
	ZoneIds []int64 `json:"plot_zone_ids,omitempty" bson:"plot_zone_ids,omitempty"`
	// Zones in the plot
	Zones []ZoneModel `json:"plot_zones" bson:"plot_zones"` 
	//created_at date time
	PlotCreatedAt time.Time `json:"plot_created_at,omitempty" bson:"plot_created_at,omitempty"`
	//saved_at date time
	PlotSavedAt time.Time `json:"plot_saved_at,omitempty" bson:"plot_saved_at,omitempty"`

	// origin of the plot: it can be bumbal, bzone or algo 
	// use the constants defined in this file to set this value for a single source of truth
	Origin string `json:"origin,omitempty" bson:"origin,omitempty"`
}

// ZoneModel struct for ZoneModel
type ZoneModel struct {
	// Unique Zone ID; use google/uuid to generate unique ids
	Id string `bson:"zone_id,omitempty" json:"zone_id"`
	//Zones Ranges
	ZoneRanges []ZoneRangeModel `bson:"zone_ranges" json:"zone_ranges"`
	//Total fuel cost
	ZoneFuelCost float64 `bson:"zone_fuel_cost,omitempty" json:"zone_fuel_cost"`
	//Total driving time
	ZoneDrivingTime int64 `bson:"zone_driving_time,omitempty" json:"zone_driving_time"`
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
	ZoneRangeId int `json:"zone_range_id,omitempty" bson:"zone_range_id"`
	// Zipcode range start
	ZipcodeFrom int `json:"zipcode_from" bson:"zipcode_from"`
	// Zipcode range end
	ZipcodeTo int `json:"zipcode_to" bson:"zipcode_to"`
	// iso country of the zone range
	IsoCountry string `json:"iso_country,omitempty" bson:"iso_country,omitempty"`
	//array of coordinates?
}

