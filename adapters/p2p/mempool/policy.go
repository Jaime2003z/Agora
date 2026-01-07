package mempool

import (
	"errors"
	"sync"
	"time"

	"github.com/Jaime2003z/Agora/adapters/p2p/protocol"
)

// ==============================
// Policy
// ==============================

type Policy struct {
	mu sync.Mutex

	lastProposal map[string]int64 // proposerID → timestamp of last proposal
	minCooldown  int64
}

func NewPolicy(minCooldown time.Duration) *Policy {
	return &Policy{
		lastProposal: make(map[string]int64),
		minCooldown:  int64(minCooldown.Seconds()),
	}
}

// ==============================
// Policy - Proposal Admission Control
// ==============================

// AllowProposal checks if a proposal should be admitted based on policy rules
func (p *Policy) AllowProposal(intent protocol.ProposalIntent) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	last, exists := p.lastProposal[intent.Payload.Proposer]
	if exists {
		if intent.Timestamp-last < p.minCooldown {
			return errors.New("mempool: proposal cooldown active")
		}
	}

	// Register timestamp
	p.lastProposal[intent.Payload.Proposer] = intent.Timestamp
	return nil
}
