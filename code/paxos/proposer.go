package paxos

import (
	"context"
	"log"
	"time"
)

type Proposer struct {
	ProposalNumber int
	Value          interface{}
}

func (p *Proposer) Propose(ctx context.Context, value interface{}, acceptors []*Acceptor) interface{} {
	// Retry logic parameters
	maxRetries := 3
	retryDelay := 50 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		select {
		case <-ctx.Done():
			log.Println("Propose cancelled or timed out:", ctx.Err())
			return nil
		default:
		}

		if i > 0 {
			log.Printf("Retry attempt %d/%d...\n", i+1, maxRetries)
			time.Sleep(retryDelay)
		}

		promises := 0
		for _, acceptor := range acceptors {
			// Simulate network delay or check context before each call if desired
			promise := acceptor.HandlePrepare(Prepare{ProposalNumber: p.ProposalNumber})
			if promise.ProposalNumber == p.ProposalNumber {
				promises++
			}
		}

		if promises > len(acceptors)/2 {
			accepted := 0
			for _, acceptor := range acceptors {
				ack := acceptor.HandleAccept(Accept{ProposalNumber: p.ProposalNumber, Value: value})
				if ack.ProposalNumber == p.ProposalNumber {
					accepted++
				}
			}

			if accepted > len(acceptors)/2 {
				return value
			}
		}

		log.Println("Consensus not reached, retrying...")
	}

	return nil
}
