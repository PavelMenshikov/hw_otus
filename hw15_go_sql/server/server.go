package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/db"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodPost {
		_, err := http.MaxBytesReader(w, r.Body, 1048576).Read(make([]byte, 1024))
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
	}
	response := fmt.Sprintf("Hello! You made a %s request to %s", r.Method, r.URL.Path)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func respondJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	respondJSON(w, users)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAllProducts()
	if err != nil {
		http.Error(w, "Error getting products", http.StatusInternalServerError)
		return
	}
	respondJSON(w, products)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id parameter is required", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	orders, err := db.GetOrdersByUser(userID)
	if err != nil {
		http.Error(w, "Error getting orders", http.StatusInternalServerError)
		return
	}
	respondJSON(w, orders)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id parameter is required", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	stats, err := db.GetUserStats(userID)
	if err != nil {
		http.Error(w, "Error getting user stats", http.StatusInternalServerError)
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
		w.Write([]byte("Online shop server is running"))
	})

	serverAddr := addr + ":" + strconv.Itoa(port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("Server is running on %s", serverAddr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
