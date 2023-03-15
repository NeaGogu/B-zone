package main

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
)

// getUserPlotIDs
//
//	@Description: function that handles the request for retrieving the plot IDs of a user
//	@receiver app
//	@param w
//	@param r
func (app *application) getUserPlotIDs(w http.ResponseWriter, r *http.Request) {

	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

	// check if we indeed have a bearer token
	if strings.Contains(jwToken, "Bearer") && jwToken[len("Bearer "):] != "" {
		jwToken = jwToken[len("Bearer "):] // remove "Bearer "
	} else {
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// parse the token to extract its information
	token, _, err := new(jwt.Parser).ParseUnverified(jwToken, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// get the user ID from the decoded claims
	for key, val := range claims {
		if key == "user_id" {
			userID, ok := val.(string)

			if ok != true {
				app.errorLog.Println("User ID is not a string")
			}

			// convert the user id from string to int as this is how it is stored in the database
			userId, err := strconv.Atoi(userID)

			if err != nil {
				app.errorLog.Println(err.Error())
			}

			// send the user id to the model which will return the user's plot IDs
			reqPlotIDs, err := app.bzoneDbModel.GetPlotIDs(userId)
			if err != nil {
				app.serverError(w, err)
			}

			// encode the output
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reqPlotIDs)

			return
		}
	}

	return
}
