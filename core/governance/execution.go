package governance

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Jaime2003z/Agora/core/ledger"
)

// ProposalStatus represents the current status of a proposal
type ProposalStatus string

const (
	StatusOpen     ProposalStatus = "open"
	StatusApproved ProposalStatus = "approved"
	StatusRejected ProposalStatus = "rejected"
)

// ExecutionResult contains the result of executing a proposal
type ExecutionResult struct {
	ProposalID string
	Status     ProposalStatus
	Reason     string
}

// CloseProposal closes a proposal and determines the result
func CloseProposal(p *Proposal, voteService interface {
	GetVotes(proposalID string) ([]Vote, error)
}) (*ExecutionResult, *Event, error) {
	if p.Closed {
		return nil, nil, errors.New("proposal is already closed")
	}

	// Get all votes for this proposal
	votes, err := voteService.GetVotes(p.ID)
	if err != nil {
		return nil, nil, err
	}

	// Calculate votes
	var yesVotes, noVotes int
	for _, vote := range votes {
		if vote.Approved {
			yesVotes++
		} else {
			noVotes++
		}
	}

	// Determine result
	var status ProposalStatus
	var eventType string

	switch {
	case yesVotes > noVotes:
		status = StatusApproved
		eventType = "ProposalAccepted"
	case noVotes > yesVotes:
		status = StatusRejected
		eventType = "ProposalRejected"
	default:
		status = StatusRejected // In case of tie, default to rejected
		eventType = "ProposalRejected"
	}

	// Create event
	event := &Event{
		Type:      eventType,
		Timestamp: time.Now().Unix(),
		Signer:    "system", // System-generated event
	}

	// Marshal the result for the event payload
	result := ExecutionResult{
		ProposalID: p.ID,
		Status:     status,
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return nil, nil, err
	}
	event.Payload = payload

	// Mark proposal as closed
	p.Closed = true

	return &result, event, nil
}

// ExecuteProposal executes the approved proposal
func ExecuteProposal(p *Proposal, chain *ledger.Chain) (*Event, error) {
	if !p.Closed {
		return nil, errors.New("cannot execute open proposal")
	}

	// In a real implementation, this would contain the actual execution logic
	// For now, we'll just create an execution event

	event := &Event{
		Type:      "ProposalExecuted",
		Timestamp: time.Now().Unix(),
		Signer:    "system",
		Payload:   []byte(`{"proposal_id":"` + p.ID + `"}`),
	}

	return event, nil
}
