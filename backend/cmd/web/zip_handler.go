package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)


func (app *application) getZipCodeCoords(w http.ResponseWriter, r *http.Request) {

	reqZipCodeFrom, err := strconv.Atoi(r.URL.Query().Get("zip_from"))
	if err != nil || reqZipCodeFrom < 0 {
		http.Error(w, "Invalid zip_from in query", http.StatusBadRequest)
		return
	}

	reqZipCodeTo, err := strconv.Atoi(r.URL.Query().Get("zip_to"))
	if err != nil || reqZipCodeTo < 0 {
		http.Error(w, "Invalid zip_to in query", http.StatusBadRequest)
		return
	}


	reqZipcodes, err := app.zipCodeDbModel.GetZipCodes(reqZipCodeFrom, reqZipCodeTo)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reqZipcodes)
	return
}
