package bumbal

import (
	"bytes"
	"bzone/backend/internal/models"
	"encoding/json"
	"net/http"
)

const (
	activitiesUrl = "activity"
)

// Output struct used for the body of the response
type Output struct {
	Result []models.ZoneModel `json:"result,omitempty"`
}

// requestBumbalActivity
//
//	@Description: makes a request to Bumbal that is meant to retrieve a user's activities
//	@param w
//	@param r
//	@return *http.Response
//	@return error
func requestBumbalActivity(w http.ResponseWriter, r *http.Request, reqBody []byte) (*http.Response, error) {
	reqUrl := baseUrl + activitiesUrl
	// set the request body so that it retrieves the address_applied field as well
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

// getZonesInfo
//
//	@Description: get the data from the request's body
//	@param r
//	@return ClustersInfo
//	@return error
func getZonesInfo(r *http.Request) (ZonesInfo, error) {
	var zonesInfo ZonesInfo
	// decode the data from the body
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&zonesInfo)
	return zonesInfo, err
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
