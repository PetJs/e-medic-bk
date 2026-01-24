// Package testutil provides test utilities.
package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestRequest creates a test HTTP request.
func TestRequest(method, path string, body interface{}) *http.Request {
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// TestRequestWithAuth creates a test HTTP request with authorization.
func TestRequestWithAuth(method, path, token string, body interface{}) *http.Request {
	req := TestRequest(method, path, body)
	req.Header.Set("Authorization", "Bearer "+token)
	return req
}

// PerformRequest performs a test request and returns the response.
func PerformRequest(t *testing.T, router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// ParseResponse parses a JSON response body.
func ParseResponse(t *testing.T, w *httptest.ResponseRecorder, v interface{}) {
	t.Helper()
	if err := json.NewDecoder(w.Body).Decode(v); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
}
