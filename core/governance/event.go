package governance

import (
	"encoding/json"
	"time"
)

// Event represents a governance event that occurs when a proposal is executed
// or when its state changes
//
// The Type field indicates the type of event (e.g., "proposal_executed", "proposal_closed")
// The Signer field indicates who triggered the event (user ID or "system")
// The Payload contains the raw event data
// The Timestamp field indicates when the event occurred
type Event struct {
	Type      string          `json:"type"`
	Signer    string          `json:"signer"`
	Payload   []byte          `json:"payload"`
	Timestamp int64           `json:"timestamp"`
	Proposal  *Proposal       `json:"proposal,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

// NewEvent creates a new governance event
func NewEvent(eventType, signer string, payload []byte) *Event {
	return &Event{
		Type:      eventType,
		Signer:    signer,
		Payload:   payload,
		Timestamp: time.Now().Unix(),
	}
}

// NewEventWithData creates a new governance event with structured data
func NewEventWithData(eventType, signer string, data interface{}) (*Event, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &Event{
		Type:      eventType,
		Signer:    signer,
		Payload:   payload,
		Timestamp: time.Now().Unix(),
	}, nil
}
