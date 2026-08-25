package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	IP       string    `json:"ip"`
	Port     int       `json:"port"`
	Status   string    `json:"status"`
	LastSeen time.Time `json:"last_seen"`
}

var (
	mu       sync.Mutex
	services = []Service{
		{
			ID:       "node-1",
			Name:     "MeshScope Local Node",
			IP:       "127.0.0.1",
			Port:     8080,
			Status:   "active",
			LastSeen: time.Now(),
		},
	}
)

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(services)
		return
	}

	if r.Method == http.MethodPost {
		var newService Service
		err := json.NewDecoder(r.Body).Decode(&newService)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		// Update existing service or add new one
		exists := false
		for i, s := range services {
			if s.ID == newService.ID {
				services[i] = newService
				exists = true
				break
			}
		}
		if !exists {
			services = append(services, newService)
		}
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newService)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/api/v1/services", servicesHandler)

	fmt.Println("🚀 MeshScope Core Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}