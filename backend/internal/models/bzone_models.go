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
	PlotId string `json:"plot_id,omitempty" bson:"plot_id,omitempty"`
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
	ZoneRangeId int64 `json:"zone_range_id,omitempty" bson:"zone_range_id"`
	// Zipcode range start
	ZipcodeFrom int64 `json:"zipcode_from" bson:"zipcode_from"`
	// Zipcode range end
	ZipcodeTo int64 `json:"zipcode_to" bson:"zipcode_to"`
	// iso country of the zone range
	IsoCountry string `json:"iso_country,omitempty" bson:"iso_country,omitempty"`
	//array of coordinates?
}

// helper method to query the database for a zone by id
// pretty much obsolete because GetZonesListById can behave in the same way
// if passed only one id
func (m *BzoneDBModel) GetZoneById(zoneId string) (*ZoneModel, error) {

	// Get the zones collection
	coll := m.DB.Collection(ZoneCollection)

	// Create a filter with the zone id
	queryFilter := bson.M{"zone_id": zoneId}

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

// helper method to query the database for a list of zones
func (m *BzoneDBModel) GetZonesListById(zoneIds ...string) ([]*ZoneModel, error) {

	// Get the zones collection
	coll := m.DB.Collection(ZoneCollection)

	queryFilter := bson.M{"zone_id": bson.M{"$in": zoneIds}}

	cur, err := coll.Find(context.TODO(), queryFilter)
	if err != nil {
		return nil, err
	}

	// don't forget to close the cursor when you're done
	defer cur.Close(context.TODO())

	var result []*ZoneModel
	if err = cur.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, ErrDocumentNotFound
	}

	return result, nil
}
