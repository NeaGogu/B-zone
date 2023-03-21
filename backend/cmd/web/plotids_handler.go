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
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// getUserPlotIDs
//
//	@Description: function that handles the request for retrieving the plot IDs of a user
//	@receiver app
//	@param w
//	@param r
func (app *application) getUserPlotIDs(w http.ResponseWriter, r *http.Request) {

	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	// check if we indeed have a bearer token
	if strings.Contains(jwToken, "Bearer") && jwToken[len("Bearer "):] != "" {
		jwToken = jwToken[len("Bearer "):] // remove "Bearer "
	} else {
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// parse the token to extract its information
	token, _, err := new(jwt.Parser).ParseUnverified(jwToken, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// get the user ID from the decoded claims
	for key, val := range claims {
		if key == "user_id" {
			userID, ok := val.(string)

			if ok != true {
				app.errorLog.Println("User ID is not a string")
			}

			// convert the user id from string to int as this is how it is stored in the database
			userId, err := strconv.Atoi(userID)

			if err != nil {
				app.errorLog.Println(err.Error())
			}

			// send the user id to the model which will return the user's plot IDs
			reqPlotIDs, err := app.bzoneDbModel.GetPlotIDs(userId)
			if err != nil {
				app.serverError(w, err)
			}

			// encode the output
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reqPlotIDs)

			return
		}
	}

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
	return
}

func (app *application) SyncBumbalZones(w http.ResponseWriter, r *http.Request) {
	// FIX: move this logic to a middleware
	t := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	token := strings.TrimPrefix(t, "Bearer ")

	// check if we indeed have a bearer token
	if token == "" {
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// parse the token to extract its information, no need to verify the token sig
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	// get the user id from the decoded claims
	uid, ok := claims["user_id"].(string)
	if !ok {
		app.serverError(w, fmt.Errorf("Could not get user id from claims"))
	}

	// convert the user id from string to int as this is how it is stored in the database
	userId, err := strconv.Atoi(uid)

	app.infoLog.Println("Syncing bumbal zones for user: ", userId)

	// get zones from Bumbal through their API
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

	// FIX: move this logic to a middleware
	t := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	token := strings.TrimPrefix(t, "Bearer ")

	// check if we indeed have a bearer token
	if token == "" {
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// parse the token to extract its information, no need to verify the token sig
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	// get the user id from the decoded claims
	uid, ok := claims["user_id"].(string)
	if !ok {
		app.serverError(w, fmt.Errorf("Could not get user id from claims"))
	}

	// convert the user id from string to int as this is how it is stored in the database
	userId, err := strconv.Atoi(uid)

	// ----------------------------
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
