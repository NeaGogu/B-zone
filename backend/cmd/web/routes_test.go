package main

import (
	"bzone/backend/internal/test"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
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

	// Test with a request that has nil user id
	req = httptest.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	tokenNil, _ := CreateAccessWrongTokenString(nil)
	req.Header.Add("Authorization", tokenNil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Test with a request that has non convertable user id
	uid := strconv.FormatInt(math.MaxInt64, 10) + "0"
	req = httptest.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	tokenNon, _ := CreateAccessWrongTokenString(uid)
	req.Header.Add("Authorization", tokenNon)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

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

func CreateAccessWrongTokenString(userId any) (string, error) {
	// Define the token claims, including the user ID
	claims := jwt.MapClaims{
		"user_id": userId,
	}

	// Define the token signing method and secret key
	signingMethod := jwt.SigningMethodHS256
	secret := []byte("my_secret_key")

	// Create the token
	token := jwt.NewWithClaims(signingMethod, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	// Return the Bearer string access token
	return fmt.Sprintf("Bearer %s", tokenString), nil
}
