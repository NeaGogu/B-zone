package bumbal

import (
	"bytes"
	kMeans "bzone/backend/cmd/k-means"
	"bzone/backend/internal/models"
	"encoding/json"
	"net/http"
)

// const baseUrl = "https://sep202302.bumbal.eu/api/v2/"
const (
	activitiesUrl = "activity"
)

// ClustersInfo struct used for retrieving the data from the body of the request
type ClustersInfo struct {
	NrClusters          int `json:"number_of_clusters,omitempty"`
	NrCandidateClusters int `json:"number_of_candidate_clusters,omitempty"`
}

// Output struct used for the body of the response
type Output struct {
	Result []models.ZoneModel `json:"result,omitempty"`
}

// RunKMeans
//
//	 @Description: the main handler, does the request to Bumbal and calls the K-Means algorithm based
//					on the input
//	 @param w
//	 @param r
func RunKMeans(w http.ResponseWriter, r *http.Request) {
	// Make the request to Bumbal
	resp, err := requestBumbalActivity(w, r)
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

		// use the collected data as input for the K-means algorithm
		computedZones, err := kMeans.KMeans(*respModel.Items, clustersInfo.NrClusters, clustersInfo.NrCandidateClusters)
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

// requestBumbalActivity
//
//	@Description: makes a request to Bumbal that is meant to retrieve a user's activities
//	@param w
//	@param r
//	@return *http.Response
//	@return error
func requestBumbalActivity(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	reqUrl := baseUrl + activitiesUrl
	// set the request body so that it retrieves the address_applied field as well
	reqBody := []byte(`{"options":{"include_address_applied":true}}`)
	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	// query BUMBAL /activity -> PUT req with TOKEN in body
	req, err := http.NewRequest(http.MethodPut, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// add the necessary headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", jwToken)

	// make the request
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

// getClustersInfo
//
//	@Description: get the data from the request's body
//	@param r
//	@return ClustersInfo
//	@return error
func getClustersInfo(r *http.Request) (ClustersInfo, error) {
	var clustersInfo ClustersInfo
	// decode the data from the body
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&clustersInfo)
	return clustersInfo, err
}

// getResponseData
//
//	@Description: get the data from the response's body and decode it as well
//	@param resp
//	@return models.ActivityListResponseBumbal
//	@return error
func getResponseData(resp *http.Response) (models.ActivityListResponseBumbal, error) {
	var respModel models.ActivityListResponseBumbal
	// decode the data from the body
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&respModel)
	return respModel, err
}
