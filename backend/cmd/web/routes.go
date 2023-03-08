package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello brothers!"))
}

// JwtChecker
//
//	@Description: a middleware that checks the JWT sent with the request
//				no JWT, request fails immediately
//	@param next
//	@return http.Handler
func JwtChecker(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		url := "https://sep202302.bumbal.eu/api/v2/authenticate/check-token"
		jwToken := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

		// query BUMBAL /check-token -> GET req with TOKEN in body
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		req.Header.Add("Authorization", jwToken)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode != http.StatusOK {
			http.Error(w, resp.Status, resp.StatusCode)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fn)
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
	router.Use(JwtChecker)

	router.Route("/test_route", func(r chi.Router) {
		r.Get("/", hello)
	})

	return router
}
