package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/db"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
	respondJSON(w, http.StatusOK, users)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAllProducts()
	if err != nil {
		http.Error(w, "Ошибка получения товаров", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, products)
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
	respondJSON(w, http.StatusOK, orders)
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
	respondJSON(w, http.StatusOK, stats)
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
	log.Printf("Запуск сервера на %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
