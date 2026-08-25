package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Service struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IP        string    `json:"ip"`
	Port      int       `json:"port"`
	Status    string    `json:"status"`
	LastSeen  time.Time `json:"last_seen"`
}

func main() {
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "online",
			"service": "meshscope-core",
		})
	})

	http.HandleFunc("/api/v1/services", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mockServices := []Service{
			{
				ID:       "srv-1",
				Name:     "MeshScope Local Node",
				IP:       "127.0.0.1",
				Port:     8080,
				Status:   "active",
				LastSeen: time.Now(),
			},
		}
		json.NewEncoder(w).Encode(mockServices)
	})

	fmt.Println("🚀 MeshScope Core Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}