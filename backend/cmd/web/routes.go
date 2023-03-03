package main

import (
	getCoordinates "bzone/backend/cmd/web/api_endpoints"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello brothers!"))
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
		r.Get("/area_points", getCoordinates.GetAreaPoint)
	})

	return router
}
