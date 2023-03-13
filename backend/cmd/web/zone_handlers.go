package main

import (
	"bzone/backend/internal/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) getZoneRanges(w http.ResponseWriter, r *http.Request) {

	// get the zone id from the url
	zoneId := chi.URLParam(r, "zone_id")
	app.infoLog.Printf("zoneId: %s", zoneId)

	zone, err := app.bzoneDbModel.GetZoneById(zoneId) 
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
	err = json.NewEncoder(w).Encode(zone)
	if err != nil {
		app.serverError(w, err)
		return
	}

	return

}
