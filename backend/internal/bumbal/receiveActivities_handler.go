package bumbal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReceiveActivities(w http.ResponseWriter, r *http.Request) {
	url := "https://sep202302.bumbal.eu/api/v2/activity"
	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	// query BUMBAL /activity -> PUT req with TOKEN in body
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", jwToken)

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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		// close response body
		defer resp.Body.Close()

		var activityResponse ActivityListResponseBumbal
		err = json.Unmarshal(body, &activityResponse)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Println(activityResponse)
		json.NewEncoder(w).Encode(activityResponse)

		return

	} else {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
}
