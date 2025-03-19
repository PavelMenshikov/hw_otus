package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен %s запрос к %s", r.Method, r.URL.Path)
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			return
		}
		log.Printf("Тело запроса: %s", string(body))
	}
	response := fmt.Sprintf("Привет! Вы сделали %s запрос к %s", r.Method, r.URL.Path)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("Ошибка записи ответа: %v", err)
	}
}

func RunServer(addr string, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RequestHandler)
	serverAddr := fmt.Sprintf("%s:%d", addr, port)
	log.Printf("Запуск сервера на %s", serverAddr)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
