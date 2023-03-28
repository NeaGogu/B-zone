package main

import (
	"bzone/backend/internal/bumbal"
	"bzone/backend/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	uidValue := r.Context().Value(ContextUserKey)
	if uidValue == nil {
		// handle the case where the uid is not present or nil
		http.Error(w, "uid not found", http.StatusInternalServerError)
		return
	}
	uid, ok := uidValue.(string)
	if !ok {
		// handle the case where the uid is not a string
		http.Error(w, "uid is not a string", http.StatusInternalServerError)
		return
	}

	// convert the user id from string to int as this is how it is stored in the database
	userId, err := strconv.Atoi(uid)
	if err != nil {
		app.serverError(w, err)
		return
	}

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

func (app *application) SyncBumbalZones(w http.ResponseWriter, r *http.Request) {

	//Get user id from the Context
	uidValue := r.Context().Value("uid")
	if uidValue == nil {
		// handle the case where the uid is not present or nil
		http.Error(w, "uid not found", http.StatusInternalServerError)
		return
	}
	uid, ok := uidValue.(string)
	if !ok {
		// handle the case where the uid is not a string
		http.Error(w, "uid is not a string", http.StatusInternalServerError)
		return
	}

	// convert the user id from string to int as this is how it is stored in the database
	userId, err := strconv.Atoi(uid)
	if err != nil {

	}

	app.infoLog.Println("Syncing bumbal zones for user: ", userId)

	// get zones from Bumbal through their API
	// Get the token
	t := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"
	token := strings.TrimPrefix(t, "Bearer ")
	bumbalZones, err := bumbal.GetZoneListReponse(token)
	if err != nil {
		// TODO: provide a better error message for the client
		app.serverError(w, err)
		return
	}

	// convert the bumbal zones to an internal plot model
	plotModel, err := bumbalZones.ToPlotModel()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// provide a name for the plot because the bumbal API does not provide one
	plotModel.Name = "Bumbal zone" + " - " + time.Now().Format("2006-01-02 15:04:05")

	//flush all the old zones with origin: bumbal
	app.infoLog.Println("Deleting old bumbal zones for user: ", userId)
	_, err = app.bzoneDbModel.DeletePlotsByOrigin(models.OriginBumbal, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// save the plot model to the database
	app.infoLog.Println("Saving new bumbal zones for user: ", userId)
	err = app.bzoneDbModel.SavePlot(userId, plotModel)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (app *application) SavePlot(w http.ResponseWriter, r *http.Request) {

	//Get user id from the Context
	uidValue := r.Context().Value(ContextUserKey)
	if uidValue == nil {
		// handle the case where the uid is not present or nil
		http.Error(w, "uid not found", http.StatusInternalServerError)
		return
	}
	uid, ok := uidValue.(string)
	if !ok {
		// handle the case where the uid is not a string
		http.Error(w, "uid is not a string", http.StatusInternalServerError)
		return
	}

	// convert the user id from string to int as this is how it is stored in the database
	userId, err := strconv.Atoi(uid)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var receivedPlot models.PlotModel

	// Try to decode the request body into the struct. If it is one of our custom error types
	// return the appropriate status code
	err = decodeJSONBody(w, r, &receivedPlot)
	if err != nil {
		var mr *malformedRequest

		// our custom error type contains
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
			return
		}

		app.errorLog.Println(err.Error())
		app.serverError(w, err)
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
	uidValue := r.Context().Value("uid")
	if uidValue == nil {
		// handle the case where the uid is not present or nil
		http.Error(w, "uid not found", http.StatusInternalServerError)
		return
	}
	uid, ok := uidValue.(string)
	if !ok {
		// handle the case where the uid is not a string
		http.Error(w, "uid is not a string", http.StatusInternalServerError)
		return
	}

	// convert the user id from string to int as this is how it is stored in the database
	userId, err := strconv.Atoi(uid)
	if err != nil {
		app.serverError(w, err)
		return
	}

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
