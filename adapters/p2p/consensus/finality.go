package consensus

import "github.com/Jaime2003z/Agora/adapters/p2p/protocol"

// ==============================
// Finality
// ==============================

type Finality struct {
	confirmations uint64
}

// confirmations = how many ticks must pass
func NewFinality(confirmations uint64) *Finality {
	return &Finality{
		confirmations: confirmations,
	}
}

// ==============================
// Finality Rules
// ==============================

func (f *Finality) CanFinalize(
	intent protocol.ProposalIntent,
	currentTick uint64,
) bool {

	// Avoid applying something too recent
	if currentTick < uint64(intent.Timestamp)+f.confirmations {
		return false
	}

	return true
}
