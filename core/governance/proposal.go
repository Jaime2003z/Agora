package governance

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

// counter is a thread-safe counter
type counter struct {
	value int64
}

// Add atomically adds delta to the counter and returns the new value
func (c *counter) Add(delta int64) int64 {
	return atomic.AddInt64(&c.value, delta)
}

// Event represents a governance event that will be written to the ledger
type Event struct {
	Type      string
	Payload   []byte
	Timestamp int64
	Signer    string
}

// Proposal represents a governance proposal
type Proposal struct {
	ID          string
	Title       string
	Description string
	CreatorID   string
	CreatedAt   int64
	Closed      bool
}

// NewProposal creates a new proposal with validation
func NewProposal(title, description, creatorID string, identityService interface{ Exists(id string) bool }) (*Proposal, *Event, error) {
	// Validate creator has a valid identity
	if !identityService.Exists(creatorID) {
		return nil, nil, ErrInvalidIdentity
	}

	now := time.Now().Unix()
	prop := &Proposal{
		ID:          generateID(),
		Title:       title,
		Description: description,
		CreatorID:   creatorID,
		CreatedAt:   now,
		Closed:      false,
	}

	// Create ProposalSubmitted event
	event := &Event{
		Type:      "ProposalSubmitted",
		Timestamp: now,
		Signer:    creatorID,
		Payload:   []byte(fmt.Sprintf(`{"id":"%s","title":"%s"}`, prop.ID, prop.Title)),
	}

	return prop, event, nil
}

// Errors
var (
	ErrInvalidIdentity = errors.New("creator does not have a valid identity")
)

// generateID creates a unique ID for proposals
func generateID() string {
	// Using a simple counter with a timestamp for uniqueness
	return fmt.Sprintf("prop_%d_%d", time.Now().Unix(), idCounter.Add(1))
}

var idCounter = &counter{value: 0}
