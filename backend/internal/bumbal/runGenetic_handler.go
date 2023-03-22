package bumbal

import (
	kMeans "bzone/backend/cmd/k-means"
	"encoding/json"
	"net/http"
)

// ZonesInfo struct used for retrieving the data from the body of the request
type ZonesInfo struct {
	NZones       int `json:"number_of_zones,omitempty"`
	NGenerations int `json:"number_of_generations,omitempty"`
}

// RunGenetic
//
//	 @Description: the main handler, does the request to Bumbal and calls the K-Means algorithm based
//					on the input
//	 @param w
//	 @param r
func RunGenetic(w http.ResponseWriter, r *http.Request) {
	// Make the request to Bumbal
	reqBody := []byte(`{"options":{"include_address_applied":true,"include_depot_address":true}}`)
	resp, err := requestBumbalActivity(w, r, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check the status codes of the response from Bumbal
	if resp.StatusCode == http.StatusMethodNotAllowed {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	} else if resp.StatusCode == http.StatusUnprocessableEntity {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	} else if resp.StatusCode == http.StatusOK {
		// If the response status is OK, get the data from the response's body
		respModel, err := getResponseData(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// get the data from the request's body
		zonesInfo, err := getZonesInfo(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// use the collected data as input for the Genetic algorithm
		computedZones, err := kMeans.KMeans(*respModel.Items, zonesInfo.NZones, zonesInfo.NGenerations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set up the response
		var output Output
		output.Result = computedZones

		// encode the response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return

	} else {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
}
