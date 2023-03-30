package bumbal

import (
	voronoi "bzone/backend/cmd/Fortunes"
	kMeans "bzone/backend/cmd/k-means"
	"encoding/json"
	"net/http"
)

func RunVoronoi(w http.ResponseWriter, r *http.Request) {
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
		//convert the output clusters to voronoi
		result, err := voronoi.VoronoiDiagram(computedClusters)

		// set up the response
		var output OutputJSON
		output.Result = *result

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
