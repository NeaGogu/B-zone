package bumbal

import (
	kMeans "bzone/backend/cmd/k-means"
	"bzone/backend/internal/models"
	"encoding/json"
	"net/http"
)

// ClustersInfo struct used for retrieving the data from the body of the request
type ClustersInfo struct {
	NrClusters          int `json:"number_of_clusters,omitempty"`
	NrCandidateClusters int `json:"number_of_candidate_clusters,omitempty"`
}

// RunKMeans
//
//	 @Description: the main handler, does the requests to Bumbal and calls the K-Means algorithm based
//					on the input
//	 @param w used to write http responses
//	 @param r used to fetch the http request
func RunKMeans(w http.ResponseWriter, r *http.Request) {
	// Collect all the activities from Bumbal
	filteredResp, err := collectAllBumbalActivities(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the data from the request's body
	var clustersInfo ClustersInfo
	clustersInfo, err = getClustersInfo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// use the collected data as input for the K-means algorithm
	var computedClusters kMeans.Clusters
	computedClusters, err = kMeans.KMeans(filteredResp, clustersInfo.NrClusters, clustersInfo.NrCandidateClusters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// turn the clusters from the k-means algorithm into zone model
	var computedZones []models.ZoneModel
	computedZones, err = kMeans.ClusterToZoneModel(computedClusters, filteredResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set up the response
	var output Output
	output.Result = computedZones
	output.ClustersResult = computedClusters

	// encode the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
