package bzone

//import (
//	swagModels "backend/internal/swag_gen"
//	"encoding/json"
//	"encoding/xml"
//	"fmt"
//	"io"
//	"net/http"
//)
//
//type Activities struct {
//	Items []swagModels.ActivityModel `json:"items,omitempty"`
//}
//
//func ReceiveActivities(w http.ResponseWriter, r *http.Request) {
//	url := "https://sep202302.bumbal.eu/api/v2/activity"
//	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"
//
//	// query BUMBAL /activity -> PUT req with TOKEN in body
//	req, err := http.NewRequest(http.MethodPut, url, nil)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	req.Header.Add("Authorization", jwToken)
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	if resp.StatusCode == http.StatusUnauthorized {
//		http.Error(w, resp.Status, resp.StatusCode)
//		return
//	} else if resp.StatusCode == http.StatusMethodNotAllowed {
//		http.Error(w, resp.Status, resp.StatusCode)
//		return
//	} else if resp.StatusCode == http.StatusUnprocessableEntity {
//		http.Error(w, resp.Status, resp.StatusCode)
//		return
//	} else if resp.StatusCode == http.StatusOK {
//		// get the data
//		body, err := io.ReadAll(resp.Body)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		// close response body
//		defer resp.Body.Close()
//
//		var activityModel Activities
//		err = xml.Unmarshal(body, &activityModel)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(activityModel)
//
//	} else {
//		http.Error(w, resp.Status, resp.StatusCode)
//		return
//	}
//}
