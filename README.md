# Paxos Distributed Consensus Lab

This project implements the Paxos consensus algorithm in Go, designed to demonstrate distributed system concepts including consensus, fault tolerance, and deployment.

## Key Features
*   **Core Paxos Logic**: Implementation of Proposer and Acceptor roles (`paxos` package).
*   **Web Service**: HTTP API to propose values and reach consensus (`cmd/server`).
*   **Local Simulation**: A standalone script to simulate Paxos locally (`cmd/simulation`).
*   **Fault Tolerance**: Proposer includes timeouts and retries using Go Context.
*   **Kubernetes Support**: Deployment configuration for containerized execution (`Deployment.yaml`).

## Project Structure
```
code/
├── cmd/
│   ├── server/      # Web service entry point
│   └── simulation/  # Local simulation entry point
├── paxos/           # Core Paxos library (Acceptor, Proposer, Messages)
├── Deployment.yaml  # Kubernetes deployment config
├── Dockerfile       # Docker build config
└── go.mod           # Go module definition
```

## How to Run

### 1. Local Simulation
```bash
go run cmd/simulation/main.go
```

### 2. Web Service
Start the server:
```bash
go run cmd/server/main.go
```
Send a proposal (in a new terminal):
```bash
curl -X POST -H "Content-Type: application/json" -d '{"ProposalNumber":1,"Value":"Test"}' http://localhost:8080/propose
```

### 3. Kubernetes
```bash
kubectl apply -f Deployment.yaml
```