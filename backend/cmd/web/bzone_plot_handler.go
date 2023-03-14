package main

import (
	"bzone/backend/internal/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

func (app *application) getBZonePlot(w http.ResponseWriter, r *http.Request) {

	reqBZonePlot, err := strconv.Atoi(r.URL.Query().Get("plot_id"))
	if err != nil || reqBZonePlot < 0 {
		http.Error(w, "Invalid plot id in query", http.StatusBadRequest)
		return
	}

	// get the plots collection
	plotsCollection := app.bzoneDbModel.DB.Collection(models.PlotCollection)

	//set query filter
	queryFilter := bson.D{{Key: "plot_id", Value: reqBZonePlot}}

	var bZonePlot models.PlotModel
	err = plotsCollection.FindOne(context.TODO(), queryFilter).Decode(&bZonePlot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bZonePlot)
	return
}
