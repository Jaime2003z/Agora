package ledger

// Event represents a ledger event in the system
type Event struct {
	// ID is the unique identifier of the event
	ID string `json:"id"`
	// Type indicates the type of the event
	Type string `json:"type"`
	// Data contains the event payload
	Data []byte `json:"data"`
	// Timestamp is when the event was created
	Timestamp int64 `json:"timestamp"`
}
