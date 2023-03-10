package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDBModel struct {
	DB *mongo.Database
}

// UserModel struct for UsersModel
type UserModel struct {
	//B-Zone User ID
	UserId int64 `json:"user_id,omitempty"`
	// unique per user
	Uuid string `json:"uuid,omitempty"`
	// plots per user
	PlotIds []int64 `json:"user_plot_ids,omitempty"`
}
