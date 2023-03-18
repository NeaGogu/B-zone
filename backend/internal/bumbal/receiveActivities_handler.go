package bumbal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://sep202302.bumbal.eu/api/v2/"
const (
	activitiesUrl = "activity"
)

func ReceiveActivities(w http.ResponseWriter, r *http.Request) {
	reqUrl := baseUrl + activitiesUrl
	reqBody := []byte(`{"options":{"include_address_applied":true}}`)
	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	fmt.Println(bytes.NewBuffer(reqBody))
	// query BUMBAL /activity -> PUT req with TOKEN in body
	req, err := http.NewRequest(http.MethodPut, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", jwToken)

	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode == http.StatusMethodNotAllowed {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	} else if resp.StatusCode == http.StatusUnprocessableEntity {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	} else if resp.StatusCode == http.StatusOK {
		// get the data
		//body, err := io.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Println(err.Error())
		//}
		//// close response body
		//defer resp.Body.Close()

		var respModel ActivityListResponseBumbal
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&respModel)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Println(respModel)
		json.NewEncoder(w).Encode(respModel)

		return

	} else {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
}
