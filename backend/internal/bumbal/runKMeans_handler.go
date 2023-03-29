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
//	 @Description: the main handler, does the request to Bumbal and calls the K-Means algorithm based
//					on the input
//	 @param w
//	 @param r
func RunKMeans(w http.ResponseWriter, r *http.Request) {
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
		clustersInfo, err := getClustersInfo(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// filter the response so that only activities that have both address applied and depot address are used
		filteredResp := filterResp(*respModel.Items)

		var computedClusters kMeans.Clusters
		computedClusters, err = kMeans.KMeans(filteredResp, clustersInfo.NrClusters, clustersInfo.NrCandidateClusters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//convert the output clusters to zones
		var computedZones []models.ZoneModel
		computedZones, err = kMeans.ClusterToZoneModel(computedClusters, filteredResp)
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
