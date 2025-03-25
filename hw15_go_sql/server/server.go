package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/db"
)

func respondJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, "Ошибка получения пользователей", http.StatusInternalServerError)
		return
	}
	respondJSON(w, users)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAllProducts()
	if err != nil {
		http.Error(w, "Ошибка получения товаров", http.StatusInternalServerError)
		return
	}
	respondJSON(w, products)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан параметр user_id", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Неверный user_id", http.StatusBadRequest)
		return
	}

	orders, err := db.GetOrdersByUser(userID)
	if err != nil {
		http.Error(w, "Ошибка получения заказов", http.StatusInternalServerError)
		return
	}
	respondJSON(w, orders)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан параметр user_id", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Неверный user_id", http.StatusBadRequest)
		return
	}

	stats, err := db.GetUserStats(userID)
	if err != nil {
		http.Error(w, "Ошибка получения статистики", http.StatusInternalServerError)
		return
	}
	respondJSON(w, stats)
}

func RunServer(addr string, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/products", productsHandler)
	mux.HandleFunc("/orders", ordersHandler)
	mux.HandleFunc("/stats", statsHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Сервер онлайн-магазина работает"))
	})

	serverAddr := addr + ":" + strconv.Itoa(port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("Запуск сервера на %s", serverAddr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
