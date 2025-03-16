package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// RequestHandler обрабатывает входящие HTTP запросы.
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен %s запрос к %s", r.Method, r.URL.Path)
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}
		log.Printf("Тело запроса: %s", string(body))
	}
	response := fmt.Sprintf("Привет! Вы сделали %s запрос к %s", r.Method, r.URL.Path)
	_, _ = w.Write([]byte(response))
}

// RunServer запускает HTTP сервер на заданном адресе и порту с установленными таймаутами.
func RunServer(addr string, port int) {
	http.HandleFunc("/", RequestHandler)
	serverAddr := fmt.Sprintf("%s:%d", addr, port)
	log.Printf("Запуск сервера на %s", serverAddr)
	srv := &http.Server{
		Addr:         serverAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
