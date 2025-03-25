package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestHandler_GET(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/test", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "Hello! You made a GET request") {
		t.Errorf("Unexpected response: %s", string(body))
	}
}

func TestRequestHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest("PUT", "http://localhost/test", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}
