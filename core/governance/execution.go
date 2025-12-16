package governance

import (
	"github.com/Jaime2003z/Agora/core/ledger"
)

type Proposal struct {
	ID          string
	Title       string
	Description string
	VotesYes    int
	VotesNo     int
	Executed    bool
}

func NewProposal(title, description string) *Proposal {
	return &Proposal{
		ID:          generateID(),
		Title:       title,
		Description: description,
	}
}

func Vote(p *Proposal, userID string, vote bool) {
	if vote {
		p.VotesYes++
	} else {
		p.VotesNo++
	}
}

func CheckQuorum(p *Proposal) bool {
	// Simple quorum check - can be modified as needed
	return p.VotesYes > p.VotesNo && (p.VotesYes+p.VotesNo) >= 1
}

func ExecuteProposal(p *Proposal, chain *ledger.Chain) {
	// Execute the proposal on the blockchain
	// This is a placeholder - implement actual execution logic
	p.Executed = true
}

func generateID() string {
	// Generate a more meaningful proposal ID with a counter
	IDCounter++
	return "prop_" + string(rune('A' + (IDCounter-1)%26)) + string(rune('0' + (IDCounter-1)/26 + 1))
}

var IDCounter int