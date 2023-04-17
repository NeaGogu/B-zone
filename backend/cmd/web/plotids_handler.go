package main

import (
	"bzone/backend/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// getUserPlotIDs
//
//	@Description: function that handles the request for retrieving the plot IDs of a user
//	@receiver app
//	@param w
//	@param r
func (app *application) getUserPlotIDs(w http.ResponseWriter, r *http.Request) {

	//Get user id from the Context
	var userId = r.Context().Value(ContextUserKey).(int)

	// send the user id to the model which will return the user's plot IDs
	reqPlotIDs, err := app.bzoneDbModel.GetPlotIDs(userId)
	if err != nil {
		if errors.Is(err, models.ErrDocumentNotFound) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
	}

	// encode the output
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reqPlotIDs)

	return
}

// Handler for getting a plot by its ID
// Reads the plot ID from the URL and returns the plot if successful
// if the plot does not exist, it returns a 404, otherwise it returns a 500 response
func (app *application) GetPlotById(w http.ResponseWriter, r *http.Request) {

	plotId := chi.URLParam(r, "plotId")
	if plotId == "" {
		http.Error(w, "Plot ID is required", http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Getting plot with id: ", plotId)

	plot, err := app.bzoneDbModel.GetPlotById(plotId)
	if err != nil {

		// return a different error if the zone with zoneId does not exist
		if errors.Is(err, models.ErrDocumentNotFound) {
			app.notFound(w)
			return
		}

		app.serverError(w, err)
		return
	}

	// encode the output
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plot)
}

func (app *application) SavePlot(w http.ResponseWriter, r *http.Request) {

	//Get user id from the Context
	var userId = r.Context().Value(ContextUserKey).(int)

	var receivedPlot models.PlotModel

	// Try to decode the request body into the struct. If it is one of our custom error types
	// return the appropriate status code
	err := decodeJSONBody(w, r, &receivedPlot)
	if err != nil {
		var mr *malformedRequest

		// our custom error type contains
		if errors.As(err, &mr) {
			err = mr
			http.Error(w, mr.msg, mr.status)
		}

		return
	}

	// generate a plot id even if it was provided by the client
	// this is to ensure that the plot id is unique
	receivedPlot.PlotId = uuid.New().String()
	now := time.Now()
	// NOTE: there is no functionality to update an existing plot CURRENTLY, so created at and saved at are the same
	receivedPlot.PlotCreatedAt = now
	receivedPlot.PlotSavedAt = now

	// NOTE: for now we assume that all plots coming to this endpoint are created by the user
	receivedPlot.Origin = models.OriginBzone

	app.infoLog.Printf("Saving plot with id %s for user %d ", receivedPlot.PlotId, userId)
	err = app.bzoneDbModel.SavePlot(userId, &receivedPlot)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (app *application) DeletePlotById(w http.ResponseWriter, r *http.Request) {

	// get the plot id from the url parameter
	plotId := chi.URLParam(r, "plotId")
	if plotId == "" {
		http.Error(w, "Plot ID is required", http.StatusBadRequest)
		return
	}

	//Get user id from the Context
	var userId = r.Context().Value(ContextUserKey).(int)

	// --------------------

	count, err := app.bzoneDbModel.DeletePlotById(plotId, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// no plot was deleted, so the plot id was not found
	if count == 0 {
		message := fmt.Sprintf("Plot with id %s not found", plotId)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
