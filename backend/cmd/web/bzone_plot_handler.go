package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// getBZonePlot
//
//	@Description: function that handles the request for retrieving the one Plot based on the plot id
//	@receiver app
//	@param w
//	@param r
func (app *application) getBZonePlot(w http.ResponseWriter, r *http.Request) {
	//receive the plot id from the query
	reqBZonePlot, err := strconv.Atoi(r.URL.Query().Get("plot_id"))
	if err != nil || reqBZonePlot < 0 {
		http.Error(w, "Invalid plot id in query", http.StatusBadRequest)
		return
	}

	// send the plot id to the model which will return the the plot model
	bZonePlot, err := app.bzoneDbModel.GetPlot(reqBZonePlot)
	if err != nil {
		http.Error(w, "Plot does not exist", http.StatusNotFound)
		return
	}

	// encode the output
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bZonePlot)
	return
}
