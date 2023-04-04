package bumbal

import (
	"bytes"
	kMeans "bzone/backend/cmd/k-means"
	"bzone/backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	activitiesUrl = "activity"
)

// Output struct used for the body of the response of the K-means algo handler
type Output struct {
	Result         []models.ZoneModel `json:"zone_model_result,omitempty"`
	ClustersResult kMeans.Clusters    `json:"clusters_result,omitempty"`
}

// requestBumbalActivity
//
//	@Description: makes a request to Bumbal that is meant to retrieve a user's activities
//	@param w used to write http responses
//	@param r used to fetch the http request
//	@return *http.Response return a http response for further processing
//	@return error for debugging purposes/robustness
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
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	return resp, err
}

// collectAllBumbalActivities
//
//	@Description: Function that first gets the number of activites stored at Bumbal and based on this information
//					it collects all activities and then filters them in order to remove depot activities
//	@param w used to write http responses
//	@param r used to fetch the http request
//	@return []models.ActivityModelBumbal returns the collected activities
//	@return error for debugging purposes/robustness
func collectAllBumbalActivities(w http.ResponseWriter, r *http.Request) ([]models.ActivityModelBumbal, error) {
	// Make a request to Bumbal to retrieve the number of activities
	reqBody := []byte(`{"count_only":true}`)
	resp, err := requestBumbalActivity(w, r, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Check the status codes of the response from Bumbal
	if resp.StatusCode != http.StatusOK {
		http.Error(w, resp.Status, resp.StatusCode)
		return nil, err
	}

	// If the response status is OK, get the data from the response's body
	respModel, err := getResponseData(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Get the total number of needed requests
	var numberOfRequests int32
	if respModel.CountFiltered%100 == 0 {
		numberOfRequests = respModel.CountFiltered / 100
	} else {
		numberOfRequests = respModel.CountFiltered/100 + 1
	}

	// Construct the layout of the requests to get activities
	reqBodyHead := []byte(`{"options":{"include_address_applied":true,"include_depot_address": true},
						"sorting_column":"id","as_list": true,"count_only":false,"limit":100,"offset":`)
	reqBodyTail := []byte(`}`)

	// Make a variable that will store a list of all activities collected from all requests
	var fullRespModel models.ActivityListResponseBumbal
	emptyList := make([]models.ActivityModelBumbal, 0)
	fullRespModel.Items = &emptyList

	// Set up a channel for collecting the responses from each request
	respChan := make(chan *http.Response, numberOfRequests)
	var wg sync.WaitGroup // wait group that waits for the goroutines to finish

	for i := 0; i < int(numberOfRequests); i++ {
		wg.Add(1)
		// go routine that will make one of the requests to Bumbal
		go func(i int) {
			defer wg.Done()
			offset := []byte(fmt.Sprintf("%d", i*100))
			reqBody = append(append(reqBodyHead, offset...), reqBodyTail...)

			partResponse, err := requestBumbalActivity(w, r, reqBody)
			if err != nil {
				return
			}

			respChan <- partResponse

			// wait 500 milliseconds in order to not get blacklisted by Bumbal
			time.Sleep(500 * time.Millisecond)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel
	close(respChan)

	// Take care of each response stored in the channel
	for partResponse := range respChan {
		// Check the status codes of the response from Bumbal
		if partResponse.StatusCode != http.StatusOK {
			http.Error(w, resp.Status, resp.StatusCode)
			return nil, err
		}
		// If the response status is OK, get the data from the response's body
		respModel, err := getResponseData(partResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, err
		}
		fullRespModelItems := *fullRespModel.Items
		respModelItems := *respModel.Items
		newItems := append(fullRespModelItems, respModelItems...)
		fullRespModel.Items = &newItems
	}

	// filter the response so that only activities that have both address applied and depot address are used
	filteredResp := filterResp(*fullRespModel.Items)

	return filteredResp, err
}

// filterResp
//
//	@Description: filters the response from Bumbal so that only activities with address and depot address are used
//	@param respModel used for parsing the unfiltered activities
//	@return []models.ActivityModelBumbal returns the filtered activities
func filterResp(respModel []models.ActivityModelBumbal) []models.ActivityModelBumbal {
	var filteredResp []models.ActivityModelBumbal
	for _, activity := range respModel {
		if activity.AddressApplied != nil && activity.DepotAddress != nil {
			filteredResp = append(filteredResp, activity)
		}
	}
	return filteredResp
}

// getClustersInfo
//
//	@Description: get the data from the request's body
//	@param r used to fetch the http request
//	@return ClustersInfo to fill in the ClustersInfo data structure
//	@return error for debugging purposes/robustness
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
//	@param r used to fetch the http request
//	@return ZonesInfo to fill in the ZonesInfo data sttucture
//	@return error for debugging purposes/robustness
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
//	@param resp http response to be processed in here
//	@return models.ActivityListResponseBumbal to store the response data
//	@return error for debugging purposes/robustness
func getResponseData(resp *http.Response) (models.ActivityListResponseBumbal, error) {
	var respModel models.ActivityListResponseBumbal
	// decode the data from the body
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&respModel)
	return respModel, err
}
