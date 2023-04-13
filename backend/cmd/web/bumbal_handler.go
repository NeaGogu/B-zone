package main

import (
	"bzone/backend/internal/bumbal"
	"bzone/backend/internal/models"
	"net/http"
	"strings"
)

func (app *application) SyncBumbalZones(w http.ResponseWriter, r *http.Request) {

	//Get user id from the Context
	var userId = r.Context().Value(ContextUserKey).(int)
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
	plotModel.Name = "Initial zone"

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
