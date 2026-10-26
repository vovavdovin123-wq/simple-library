package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type infoResponse struct {
	Version    string `json:"version"`
	ServerTime string `json:"server_time"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := infoResponse{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("ошибка ответа:", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "This page does not exist")
}

func main() {
	http.HandleFunc("/api/v1/info", infoHandler)
	http.HandleFunc("/notfound", notFoundHandler)

	addr := ":8080"
	fmt.Println("Сервер запущен на http://localhost" + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
