package governance

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

// Vote represents a vote on a proposal
type Vote struct {
	ID         string
	ProposalID string
	VoterID    string
	Approved   bool
	Timestamp  int64
}

// VoteService defines the interface for vote-related operations
type VoteService interface {
	HasVoted(proposalID, voterID string) (bool, error)
	AddVote(vote Vote) error
	GetVotes(proposalID string) ([]Vote, error)
}

// CastVote casts a vote on a proposal
func CastVote(proposal *Proposal, voterID string, approved bool, voteService VoteService, identityService interface{ Exists(id string) bool }) (*Vote, *Event, error) {
	// Validate proposal is open
	if proposal.Closed {
		return nil, nil, errors.New("cannot vote on closed proposal")
	}

	// Validate voter identity
	if !identityService.Exists(voterID) {
		return nil, nil, errors.New("voter does not have a valid identity")
	}

	// Check if voter has already voted
	hasVoted, err := voteService.HasVoted(proposal.ID, voterID)
	if err != nil {
		return nil, nil, err
	}
	if hasVoted {
		return nil, nil, errors.New("voter has already voted on this proposal")
	}

	// Create vote
	now := time.Now().Unix()
	vote := Vote{
		ID:         generateVoteID(),
		ProposalID: proposal.ID,
		VoterID:    voterID,
		Approved:   approved,
		Timestamp:  now,
	}

	// Add vote to storage
	if err := voteService.AddVote(vote); err != nil {
		return nil, nil, err
	}

	// Create VoteCast event
	event := &Event{
		Type:      "VoteCast",
		Timestamp: now,
		Signer:    voterID,
	}

	// Include vote details in the payload
	voteData := struct {
		ProposalID string `json:"proposal_id"`
		VoteID     string `json:"vote_id"`
		Approved   bool   `json:"approved"`
	}{
		ProposalID: proposal.ID,
		VoteID:     vote.ID,
		Approved:   approved,
	}

	payload, err := json.Marshal(voteData)
	if err != nil {
		return nil, nil, err
	}
	event.Payload = payload

	return &vote, event, nil
}

// counter is a simple thread-safe counter
type counter struct {
	value int64
}

// Add increments the counter and returns the new value
func (c *counter) Add(delta int64) int64 {
	return atomic.AddInt64(&c.value, delta)
}

// generateVoteID creates a unique ID for votes
func generateVoteID() string {
	// Using a simple counter with a timestamp for uniqueness
	return fmt.Sprintf("vote_%d_%d", time.Now().Unix(), voteCounter.Add(1))
}

var voteCounter = &counter{value: 0}
