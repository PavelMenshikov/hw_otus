package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler_GET(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/users", nil)
	w := httptest.NewRecorder()
	getUsers(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestRequestHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest("PUT", "http://localhost/users", nil)
	w := httptest.NewRecorder()
	getUsers(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}
