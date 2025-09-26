package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestData struct {
	Number int `json:"num"`
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Number < 0 || data.Number > 99 {
		http.Error(w, "Number must be between 0 and 99", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var resp map[string]interface{}

	target := rand.Intn(100)
	fmt.Printf("Received guess: %d, Target: %d\n", data.Number, target)
	if target == data.Number {
		resp = map[string]interface{}{
			"message": "Congratulations",
			"target":  target,
		}
	} else {
		resp = map[string]interface{}{
			"message": "Try Again",
			"target":  target,
		}
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)
	r.Post("/submit", submitHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
