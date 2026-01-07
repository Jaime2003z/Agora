package mempool

import "github.com/Jaime2003z/Agora/adapters/p2p/protocol"

// ==============================
// Intents (eventos canónicos)
// ==============================

type VoteIntent struct {
	MessageID string
	ProjectID string
	VoterID   string
	Approve   bool
	Weight    float32
	Timestamp int64
	SenderID  string
}

// ==============================
// Input Interfaces
// ==============================

// ProposalSink is used by protocol handlers
type ProposalSink interface {
	AddProposal(intent protocol.ProposalIntent) error
}

// VoteSink (will come later)
type VoteSink interface {
	AddVote(intent VoteIntent) error
}

// ==============================
// Mempool público
// ==============================

type Mempool struct {
	store  *Store
	policy *Policy
}

func New(store *Store, policy *Policy) *Mempool {
	return &Mempool{
		store:  store,
		policy: policy,
	}
}

// ==============================
// Proposal Ingestion
// ==============================

// Store returns the underlying store of the mempool
func (m *Mempool) Store() *Store {
	return m.store
}

func (m *Mempool) AddProposal(intent protocol.ProposalIntent) error {

	// 1️⃣ Policy (anti-spam, limits, cooldown)
	if err := m.policy.AllowProposal(intent); err != nil {
		return err
	}

	// 2️⃣ Deduplication + storage
	return m.store.AddProposal(intent)
}
