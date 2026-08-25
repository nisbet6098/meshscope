package main

import (
	"encoding/json"
	"fmt"
	"net"
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

func main() {
	fmt.Println("🛰️  MeshScope Discovery Agent active...")

	portsToScan := map[int]string{
		80:   "HTTP Web Server",
		443:  "HTTPS Secure Server",
		22:   "SSH Service",
		8080: "MeshScope Core API",
		3000: "Node/React App",
		5432: "PostgreSQL Database",
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

				jsonData, _ := json.Marshal(payload)
				fmt.Printf(" Found active service on port %d! Payload: %s\n", port, string(jsonData))
			}
		}

		time.Sleep(10 * time.Second)
	}
}