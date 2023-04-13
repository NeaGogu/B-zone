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
	type TestObject struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		body     string
		expected interface{}
		wantErr  bool
	}{
		{
			name: "valid JSON body",
			body: `{"ID": 123, "Name": "test"}`,
			expected: &TestObject{
				ID:   123,
				Name: "test",
			},
			wantErr: false,
		},

		{
			name:     "request body too large",
			body:     strings.Repeat("a", 1048577),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "unknown field in JSON",
			body:     `{"ID": 123, "Name": "test", "InvalidField": "invalid"}`,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "invalid value for field in JSON",
			body:     `{"ID": "invalid", "Name": "test"}`,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "empty request body",
			body:     "",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "multiple JSON objects in request body",
			body:     `{"ID": 123, "Name": "test"}{"ID": 456, "Name": "test2"}`,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			var dst TestObject
			err := decodeJSONBody(w, req, &dst)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !tt.wantErr {
				if !jsonEqual(&dst, tt.expected) {
					t.Fatalf("Unexpected result. Got %v, expected %v", dst, tt.expected)
				}
			}
		})
	}

	t.Run("Wrong content type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"ID": 123, "Name": "test"}`))
		req.Header.Set("Content-Type", "text/plain")

		w := httptest.NewRecorder()

		var dst TestObject
		err := decodeJSONBody(w, req, &dst)

		if (err != nil) != true {
			t.Fatalf("Unexpected error: %v", err)
		}

	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		// Create a request with malformed JSON payload
		reqBody := bytes.NewBufferString(`{"name": "Alice", "age": 30`)
		req, err := http.NewRequest("POST", "/", reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Call the decodeJSONBody function
		var dst interface{}
		err = decodeJSONBody(httptest.NewRecorder(), req, &dst)

		if (err != nil) != true {
			t.Fatalf("Unexpected error: %v", err)
		}

	})
}

// jsonEqual is a helper function to compare two JSON-encoded objects for equality
func jsonEqual(a, b interface{}) bool {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return false
	}

	bJSON, err := json.Marshal(b)
	if err != nil {
		return false
	}

	return bytes.Equal(aJSON, bJSON)
}

func TestError(t *testing.T) {
	mr := malformedRequest{
		status: 0,
		msg:    "",
	}

	mr.Error()

}
