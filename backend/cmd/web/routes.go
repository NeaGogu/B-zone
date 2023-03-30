package main

import (
	"bzone/backend/internal/bumbal"
	"context"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello brothers!"))
}

// Context key for the token check
type ContextKey string

const ContextUserKey ContextKey = "user_id"

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

func JWTRequestChecker(next http.Handler) http.Handler {
	//Checks if the token was sent correctly
	fn := func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization") // -> "Bearer <TOKEN>"

		jwToken := strings.TrimPrefix(t, "Bearer ")

		// check if we indeed have a bearer token
		if jwToken == "" {
			http.Error(w, "There is a problem with your token", http.StatusBadRequest)
			return
		}

		// parse the token to extract its information, no need to verify the token sig
		parsedToken, _, err := new(jwt.Parser).ParseUnverified(jwToken, jwt.MapClaims{})
		if err != nil {
			http.Error(w, "Cannot Parse the token", http.StatusInternalServerError)
			return
		}

		claims := parsedToken.Claims.(jwt.MapClaims)
		// get the user id from the decoded claims
		uid, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "Cannot get user_id from decoded claims", http.StatusInternalServerError)

			return
		} else if uid == "" {
			http.Error(w, "Missing user ID in your token", http.StatusInternalServerError)
			return
		}

		// convert the user id from string to int as this is how it is stored in the database
		userId, err := strconv.Atoi(uid)
		if err != nil {
			// userId is nil
			http.Error(w, "uid is not a string", http.StatusInternalServerError)
			return
		}

		// Store the uid in the request context
		ctx := context.WithValue(r.Context(), ContextUserKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))

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
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*", "*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	app.infoLog.Println("CORS enabled!")

	// for testing purposes, does not require JWT authorization
	// should not be used in production
	router.Route("/test", func(r chi.Router) {
		r.Get("/", hello)

		r.Route("/zip", func(r chi.Router) {

			r.Get("/coordinates", app.getZipCodeCoords)
		})
		r.Route("/plot", func(r chi.Router) {
			r.Get("/{plotId}", app.GetPlotById)
			r.Put("/sync", app.SyncBumbalZones)
			r.Post("/save", app.SavePlot)
			r.Delete("/{plotId}", app.DeletePlotById)
		})

	})

	// authorized routes
	router.Group(func(r chi.Router) {
		r.Use(JwtChecker)
		r.Use(JWTRequestChecker)

		r.Route("/zip", func(r chi.Router) {
			r.Get("/coordinates", app.getZipCodeCoords)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/plots", app.getUserPlotIDs)
		})

		r.Route("/plot", func(r chi.Router) {
			r.Get("/{plotId}", app.GetPlotById)
			r.Put("/sync", app.SyncBumbalZones)
			r.Post("/save", app.SavePlot)
			r.Delete("/{plotId}", app.DeletePlotById)
		})

		r.Route("/bumbal", func(r chi.Router) {
			r.Put("/algorithm/kmeans", bumbal.RunKMeans)
			r.Put("/algorithm/genetic", bumbal.RunGenetic)
			// r.Put("/algorithm/voronoi", bumbal.RunVoronoi)
		})

	})

	// log all the routes mounted on the router
	app.infoLog.Println("Mounted routes:")
	err := chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		app.infoLog.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
	if err != nil {
		panic(err)
	}

	return router
}
