package main

import (
	_ "bzone/backend/cmd/web/api_endpoints"
	"bzone/backend/internal/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello brothers!"))
}

// TODO move this to a separate file
func (app *application) getZipCodeCoords(w http.ResponseWriter, r *http.Request) {
	// TODO read from body
	reqZipCode := r.URL.Query().Get("zip_code")
	queryFilter := bson.D{{"code", reqZipCode}}
	coll := app.zipCodeDbModel.DB.Collection("coordinates")

	var zipCode models.ZipCode
	err := coll.FindOne(context.TODO(), queryFilter).Decode(&zipCode)
	if err != nil {
		http.Error(w, "Could not get zipcode", http.StatusInternalServerError)
		return
	}
	 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(zipCode)
	return
}



func (app *application) routes() http.Handler {
	// init the router
	router := chi.NewRouter()

	// TODO
	// router.NotFound()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/test_route", func(r chi.Router) {
		r.Get("/", hello)
	})

	router.Route("/zip", func(r chi.Router) {
		r.Get("/area_points", app.getZipCodeCoords)
	})

	return router
}
