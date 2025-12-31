package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"paxos-lab/paxos"
	"sync"
	"time"
)

var (
	acceptors = []*paxos.Acceptor{{}, {}, {}}
	mu        sync.Mutex
)

func proposeHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ProposalNumber int
		Value          string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	proposer := paxos.Proposer{ProposalNumber: body.ProposalNumber, Value: body.Value}
	mu.Lock()
	// Create local references to acceptors for the simulation
	localAcceptors := make([]*paxos.Acceptor, len(acceptors))
	copy(localAcceptors, acceptors)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	value := proposer.Propose(ctx, body.Value, localAcceptors)
	mu.Unlock()

	if value != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Consensus reached: %s\n", value)
	} else {
		w.WriteHeader(http.StatusConflict) // Using 409 Conflict for consensus failure
		fmt.Fprintf(w, "Consensus not reached\n")
	}
}

func main() {
	http.HandleFunc("/propose", proposeHandler)
	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
