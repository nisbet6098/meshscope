package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type ServicePayload struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	IP       string    `json:"ip"`
	Port     int       `json:"port"`
	Status   string    `json:"status"`
	LastSeen time.Time `json:"last_seen"`
}

func checkPort(port int) bool {
	timeout := 500 * time.Millisecond
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func registerService(payload ServicePayload) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	resp, err := http.Post("http://localhost:8080/api/v1/services", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("⚠️ Server unreachable for port %d\n", payload.Port)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		fmt.Printf("✅ Registered active service: %s (Port %d)\n", payload.Name, payload.Port)
	}
}

func main() {
	fmt.Println("🛰️  MeshScope Agent: Dynamic Service Registration active...")

	portsToScan := map[int]string{
		80:   "HTTP Web Server",
		443:  "HTTPS Secure Server",
		22:   "SSH Service",
		8080: "MeshScope Core API",
		5173: "Vite React Dashboard",
		3000: "Node Application",
	}

	for {
		for port, serviceName := range portsToScan {
			if checkPort(port) {
				payload := ServicePayload{
					ID:       fmt.Sprintf("agent-discovered-%d", port),
					Name:     serviceName,
					IP:       "127.0.0.1",
					Port:     port,
					Status:   "active",
					LastSeen: time.Now(),
				}
				registerService(payload)
			}
		}
		time.Sleep(10 * time.Second)
	}
}