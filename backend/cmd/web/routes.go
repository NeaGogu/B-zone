package main

import (
	_ "bzone/backend/cmd/web/api_endpoints"
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

		if resp.StatusCode == http.StatusUnauthorized {
			http.Error(w, resp.Status, resp.StatusCode)
			return
		} else if resp.StatusCode == http.StatusMethodNotAllowed {
			http.Error(w, resp.Status, resp.StatusCode)
			return
		} else if resp.StatusCode == http.StatusUnprocessableEntity {
			http.Error(w, resp.Status, resp.StatusCode)
			return
		} else if resp.StatusCode == http.StatusOK {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, resp.Status, resp.StatusCode)
			return
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

	// for testing purposes, does not require JWT authorization
	router.Route("/test", func(r chi.Router) {
		r.Get("/", hello)
		r.Route("/zip", func(r chi.Router) {

			r.Get("/coordinates", app.getZipCodeCoords)
		})
	})

	router.Group(func(r chi.Router) {
		r.Use(JwtChecker)

		r.Route("/zip", func(r chi.Router) {
			r.Get("/coordinates", app.getZipCodeCoords)
		})
	})

	return router
}
