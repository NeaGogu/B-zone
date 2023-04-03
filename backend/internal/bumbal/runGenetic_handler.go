package bumbal

import (
	"bzone/backend/cmd/genetic"
	"bzone/backend/internal/models"
	"encoding/json"
	"net/http"
	"time"
)

// ZonesInfo struct used for retrieving the data from the body of the request
type ZonesInfo struct {
	NZones       int           `json:"number_of_zones,omitempty"`
	NGenerations int           `json:"number_of_generations,omitempty"`
	MaxDuration  time.Duration `json:"maximum_runtime,omitempty"`
}

// RunGenetic
//
//	 @Description: the main handler, does the requests to Bumbal and calls the Genetic algorithm based
//					on the input
//	 @param w
//	 @param r
func RunGenetic(w http.ResponseWriter, r *http.Request) {
	// Collect all the activities from Bumbal
	filteredResp, err := collectAllBumbalActivities(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the data from the request's body
	var zonesInfo ZonesInfo
	zonesInfo, err = getZonesInfo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// use the collected data as input for the Genetic algorithm
	var computedZones []models.ZoneModel
	computedZones, err = genetic.RunGeneticAlgorithm(filteredResp, zonesInfo.NZones, zonesInfo.NGenerations,
		zonesInfo.MaxDuration*time.Minute)
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
}
