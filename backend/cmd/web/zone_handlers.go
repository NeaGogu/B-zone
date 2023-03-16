package main

import (
	"bzone/backend/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

type zoneRangesRequestBody struct {
	ZoneIds []string `json:"zone_ids"`
}

// getZoneRanges returns the zone ranges for the given list of zone ids
// the zone ids must be present in the request body as a struct
// of type zoneRangeRequestBody
func (app *application) getZoneRanges(w http.ResponseWriter, r *http.Request) {

	// get the list of zoneIds from the request body
	var reqBody zoneRangesRequestBody

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// zones, err := app.bzoneDbModel.GetZoneById(zoneId)
	zones, err := app.bzoneDbModel.GetZonesListById(reqBody.ZoneIds...)
	if err != nil {

		// return a different error if the zone with zoneId does not exist
		if errors.Is(err, models.ErrDocumentNotFound) {
			app.notFound(w)
			return
		}

		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(zones)
	if err != nil {
		app.serverError(w, err)
		return
	}

	return

}
