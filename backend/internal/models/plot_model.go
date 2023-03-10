package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PlotDBModel struct {
	DB *mongo.Database
}

// PlotModel struct for ZonesPLotModel
type PlotModel struct {
	// Unique Plot ID
	PlotId int64 `json:"plot_id,omitempty"`
	// Plot Name
	Name string `json:"plot_name,omitempty"`
	//Zones in the plot
	ZoneIds []int64 `json:"plot_zone_ids,omitempty"`
	//created_at date time
	PlotCreatedAt time.Time `json:"plot_created_at,omitempty"`
	//saved_at date time
	PlotSavedAt time.Time `json:"plot_saved_at,omitempty"`
}
