# 🚀 gRPC Microservices Demo (Go)

A minimal distributed system built with **Go + gRPC + Protocol Buffers** to understand microservices communication from first principles.

This project contains three independent services:

- **User Service** – Returns user details  
- **Inventory Service** – Checks product stock  
- **Order Service** – Orchestrates calls to User & Inventory services to create an order  

Each service:
- Runs on its own port  
- Has its own gRPC server  
- Communicates via strongly-typed Protobuf contracts  
- Demonstrates real service-to-service RPC calls  

---

## 🧠 What This Project Demonstrates

- gRPC server & client implementation in Go  
- Protobuf contract definition (`.proto`)  
- Service-to-service communication  
- Runtime reflection for debugging (`grpcurl` support)  
- Basic distributed system architecture  

---

## 🔧 Tech Stack

- Go  
- gRPC  
- Protocol Buffers  

---

## 🏗 Architecture

Client → Order Service
↓
User Service
↓
Inventory Service


Each service runs independently and communicates over gRPC using HTTP/2 and binary Protocol Buffers.