package main

import (
	"bzone/backend/internal/models"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func (app *application) getZipCodeCoords(w http.ResponseWriter, r *http.Request) {
	// TODO read from body
	reqZipCode := r.URL.Query().Get("zip_code")
	if reqZipCode == ""{
		http.Error(w, "Missing zip code in query", http.StatusBadRequest)
		return
	}

	queryFilter := bson.D{{Key: "code", Value: reqZipCode}}

	// TODO add this collection as a constant
	coll := app.zipCodeDbModel.DB.Collection("coordinates")

	var zipCode models.ZipCode
	err := coll.FindOne(context.TODO(), queryFilter).Decode(&zipCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(zipCode)
	return
}
