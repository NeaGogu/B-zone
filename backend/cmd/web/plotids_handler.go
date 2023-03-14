package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
)

// getUserPlotIDs
//
//	@Description: function that handles the request for retrieving the plot IDs of a user
//	@receiver app
//	@param w
//	@param r
func (app *application) getUserPlotIDs(w http.ResponseWriter, r *http.Request) {

	jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"
	jwToken = jwToken[len("Bearer "):]       // remove "Bearer "

	// parse the token to extract its information
	token, _, err := new(jwt.Parser).ParseUnverified(jwToken, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "There is a problem with your token", http.StatusBadRequest)
		return
	}

	// get the user ID from the decoded claims
	for key, val := range claims {
		if key == "user_id" {
			userID, ok := val.(string)

			if ok != true {
				fmt.Println("User ID is not a string")
			}

			// convert the user id from string to int as this is how it is stored in the database
			userId, err := strconv.Atoi(userID)

			if err != nil {
				fmt.Println(err.Error())
			}

			// send the user id to the model which will return the user's plot IDs
			reqPlotIDs, err := app.bzoneDbModel.GetPlotIDs(userId)

			if err != nil {
				fmt.Println(err.Error())
			}

			// encode the output
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reqPlotIDs)

			return
		}
	}

	return
}
