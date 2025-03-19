package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestHandler_GET(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус OK, получен %v", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "GET") {
		t.Errorf("Ответ не содержит ожидаемого текста GET, получено: %s", string(body))
	}
}

func TestRequestHandler_POST(t *testing.T) {
	data := "тестовые данные"
	req := httptest.NewRequest("POST", "http://localhost/test", strings.NewReader(data))
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус OK, получен %v", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "POST") {
		t.Errorf("Ответ не содержит ожидаемого текста POST, получено: %s", string(body))
	}
}
