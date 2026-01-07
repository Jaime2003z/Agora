package mempool

import (
	"errors"
	"sync"

	"github.com/Jaime2003z/Agora/adapters/p2p/protocol"
)

// ==============================
// Internal Store
// ==============================

type Store struct {
	mu sync.RWMutex

	proposals map[string]protocol.ProposalIntent
	order     []string // ordered MessageIDs
}

func NewStore() *Store {
	return &Store{
		proposals: make(map[string]protocol.ProposalIntent),
		order:     []string{},
	}
}

// ==============================
// Proposals
// ==============================

func (s *Store) AddProposal(intent protocol.ProposalIntent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Deduplication: reject proposals with duplicate MessageID
	if _, exists := s.proposals[intent.MessageID]; exists {
		return errors.New("mempool: duplicate proposal intent")
	}

	s.proposals[intent.MessageID] = intent
	s.order = append(s.order, intent.MessageID)

	return nil
}

// ListProposals returns proposals in arrival order
func (s *Store) ListProposals() []protocol.ProposalIntent {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]protocol.ProposalIntent, 0, len(s.order))
	for _, id := range s.order {
		result = append(result, s.proposals[id])
	}
	return result
}

// RemoveProposal removes a proposal from the store after consensus commit
func (s *Store) RemoveProposal(messageID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.proposals, messageID)

	for i, id := range s.order {
		if id == messageID {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}
}
