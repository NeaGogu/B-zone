package main

import (
	"encoding/json"
	"net/http"
)


func (app *application) getZipCodeCoords(w http.ResponseWriter, r *http.Request) {

	reqZipCodeFrom := r.URL.Query().Get("zip_from")
	reqZipCodeTo := r.URL.Query().Get("zip_to")
	if reqZipCodeFrom== ""{
		http.Error(w, "Missing zip_from in query", http.StatusBadRequest)
		return
	}

	if reqZipCodeTo== ""{
		http.Error(w, "Missing zip_to in query", http.StatusBadRequest)
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
