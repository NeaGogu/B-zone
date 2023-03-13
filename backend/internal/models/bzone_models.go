package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	PlotId int64 `bson:"plot_id,omitempty"`
	// Plot Name
	Name string `bson:"plot_name,omitempty"`
	//Zones in the plot
	ZoneIds []int64 `bson:"plot_zone_ids,omitempty"`
	//created_at date time
	PlotCreatedAt time.Time `bson:"plot_created_at,omitempty"`
	//saved_at date time
	PlotSavedAt time.Time `bson:"plot_saved_at,omitempty"`
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


// UserModel struct for UsersModel
type UserModel struct {
	//B-Zone User ID
	UserId int64 `bson:"user_id,omitempty"`
	// unique per user
	Uuid string `bson:"uuid,omitempty"`
	// plots per user
	PlotIds []int64 `bson:"user_plot_ids,omitempty"`
}


// ZoneRangeModel struct for ZoneRangeModel
type ZoneRangeModel struct {
	// Unique Zone type ID
	ZoneRangeId int64 `bson:"zone_range_id,omitempty"`
	// Zipcode range start
	ZipcodeFrom int64 `bson:"zipcode_from"`
	// Zipcode range end
	ZipcodeTo int64 `bson:"zipcode_to"`
	// iso country of the zone range
	IsoCountry string `bson:"iso_country,omitempty"`
	//array of coordinates?
}

// helper method to query the database for a zone by id
func (m *BzoneDBModel) GetZoneById(zoneId string) (*ZoneModel, error) {

	// Get the zones collection
	coll := m.DB.Collection(ZoneCollection)

	// Create a filter with the zone id
	queryFilter := bson.M {"zone_id": zoneId} 

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







