package main

import (
	"bzone/backend/internal/test"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	_ "bzone/backend/internal/bumbal"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hello)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hello brothers!", rr.Body.String())
}

func TestJwtChecker(t *testing.T) {
	router := chi.NewRouter()

	// Set up a simple test endpoint that returns a 200 status code
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Wrap the test endpoint with JwtChecker middleware
	router.Group(func(r chi.Router) {
		r.Use(JwtChecker)

		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	// Test with a request that has no JWT
	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// Test with a request that has an invalid JWT
	req, err = http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer invalid-token")

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// Test with a request that has a valid JWT
	req = httptest.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	token, _ := test.CreateAccessTokenString("1")
	req.Header.Add("Authorization", token)

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, 401, rr.Code) //Needs Review
}

func TestJWTRequestChecker(t *testing.T) {
	router := chi.NewRouter()

	// Set up a simple test endpoint that returns a 200 status code
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Wrap the test endpoint with JwtChecker middleware
	router.Group(func(r chi.Router) {
		r.Use(JWTRequestChecker)

		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	// Test with a request that has no JWT
	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Test with a request that has an invalid JWT
	req, err = http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer ")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Test with a request that has empty user id
	req = httptest.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	token, _ := test.CreateAccessTokenString("")
	req.Header.Add("Authorization", token)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Test with a request that has valid user id
	req = httptest.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	tokenValid, _ := test.CreateAccessTokenString("32")
	req.Header.Add("Authorization", tokenValid)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRoutes(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		zipCodeDbModel: nil,
		bzoneDbModel:   nil,
	}
	handler := app.routes()

	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
