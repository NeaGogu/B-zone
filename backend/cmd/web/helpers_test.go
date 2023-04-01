package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeJSONBody(t *testing.T) {
	t.Run("valid request body", func(t *testing.T) {
		// create a request with a valid JSON body
		body := map[string]string{
			"plot_name": "MyPlot",
			"plot_id":   "30",
		}
		reqBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/plots", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// decode the request body
		var dst map[string]string
		err = decodeJSONBody(httptest.NewRecorder(), req, &dst)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if dst["plot_name"] != "MyPlot" {
			t.Errorf("unexpected value for plot_name field: %s", dst["plot_name"])
		}
		if dst["plot_id"] != "30" {
			t.Errorf("unexpected value for age field: %s", dst["plot_id"])
		}
	})

	t.Run("invalid content type", func(t *testing.T) {
		// create a request with an invalid content type header
		reqBody := strings.NewReader(`{"plot_name": "MyPlot", "plot_id": 30}`)
		req := httptest.NewRequest(http.MethodPost, "/plots", reqBody)
		req.Header.Set("Content-Type", "text/plain")

		// decode the request body
		err := decodeJSONBody(httptest.NewRecorder(), req, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
		mr, ok := err.(*malformedRequest)
		if !ok {
			t.Errorf("unexpected error type: %T", err)
		}
		if mr.status != http.StatusUnsupportedMediaType {
			t.Errorf("unexpected status code: %d", mr.status)
		}
		if mr.msg != "Content-Type header is not application/json" {
			t.Errorf("unexpected error message: %s", mr.msg)
		}
	})

	t.Run("badly-formed JSON", func(t *testing.T) {
		// create a request with a badly-formed JSON body
		reqBody := strings.NewReader(`{"plot_name": "MyPlot", "plot_id":`)
		req := httptest.NewRequest(http.MethodPost, "/plots", reqBody)
		req.Header.Set("Content-Type", "application/json")

		// decode the request body
		err := decodeJSONBody(httptest.NewRecorder(), req, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
		mr, ok := err.(*malformedRequest)
		if !ok {
			t.Errorf("unexpected error type: %T", err)
		}
		if mr.status != http.StatusBadRequest {
			t.Errorf("unexpected status code: %d", mr.status)
		}
	})

}

//assert.Equal(t, 500, responseMissingToken.StatusCode)
