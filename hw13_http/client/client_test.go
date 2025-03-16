package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunClient_GET(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Ожидался GET-запрос, но получен %s", r.Method)
		}
		if _, err := w.Write([]byte("Ответ GET")); err != nil {
			t.Errorf("Ошибка при записи ответа: %v", err)
		}
	}))
	defer server.Close()

	RunClient("GET", server.URL, "")
}

func TestRunClient_POST(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Ожидался POST-запрос, но получен %s", r.Method)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Ошибка чтения тела запроса: %v", err)
		}
		defer r.Body.Close()

		if string(body) != "тестовые данные" {
			t.Errorf("Ожидалось тело 'тестовые данные', но получено '%s'", string(body))
		}

		if _, err := w.Write([]byte("Ответ POST")); err != nil {
			t.Errorf("Ошибка при записи ответа: %v", err)
		}
	}))
	defer server.Close()

	RunClient("POST", server.URL, "тестовые данные")
}
