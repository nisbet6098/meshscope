MeshScope
Zero-touch, offline-first local service discovery and auto-routing dashboard for Homelabs, Edge Nodes, and SMB Infrastructure.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Status](https://img.shields.io/badge/status-planning-orange.svg)

---

## 🔍 What is MeshScope?

MeshScope solves the constant pain of **Config & Discovery Drift**. 

When spinning up new VMs, LXCs, or Docker containers, admins often spend 20+ minutes manually updating IP/port mappings across reverse proxies, static dashboards, and monitoring tools. MeshScope runs local micro-agents across your hypervisors to automatically detect active services, map friendly routes, and display real-time service health on an offline-first dashboard.

---

## ✨ Key Features
## 🛠️ Architecture

```text
[ Proxmox Host / LXCs ]       [ Docker Engine ]       [ Bare-Metal OS ]
          │                           │                       │
          ├──► (MeshScope Agent) ─────┼──► (MeshScope Agent) ─┘
                                      │
                                      ▼
                        [ MeshScope Core Daemon ]
                         ├── SQLite Database
                         ├── Internal Router/Proxy
                         └── Local Web UI

---                         

## 📂 Repository Structure

├── agent/            # Go micro-agent for local daemon scanning
├── server/           # Core API engine, SQLite storage, and router
├── ui/               # Web frontend dashboard (React/Svelte + Tailwind)
├── docs/             # Architecture, API specifications, and guides
└── docker-compose.yml# Quick-start local deployment stack

---

## 🚦 Getting Started (Development Setup)
Prerequisites
Go 1.22+ / Node.js 20+

- Docker & Docker Compose

- Proxmox VE (optional for full hypervisor testing)

---

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.