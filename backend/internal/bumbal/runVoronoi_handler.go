package bumbal

import (
	voronoi "bzone/backend/cmd/Fortunes"
	kMeans "bzone/backend/cmd/k-means"
	"encoding/json"
	geojson "github.com/paulmach/go.geojson"
	"net/http"
)

// RunVoronoi
//
//	 @Description: the main handler, does the requests to Bumbal and calls the K-Means algorithm based
//					on the input and then runs the Voronoi diagram
//	 @param w
//	 @param r
func RunVoronoi(w http.ResponseWriter, r *http.Request) {
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

	var output *geojson.FeatureCollection
	output, err = voronoi.VoronoiDiagram(computedClusters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
