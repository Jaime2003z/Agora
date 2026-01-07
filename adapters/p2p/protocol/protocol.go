package protocol

// ==============================
// Protocol identity & versioning
// ==============================

const (
	Name       = "agora"
	Version    = "1.0.0"
	ProtocolID = "/" + Name + "/" + Version
)

// ==============================
// Base protocol envelope
// ==============================

// Envelope is the outer container for all protocol messages.
// The semantic meaning is defined by Message.Type.
type Envelope struct {
	Message Message `json:"message"`
}
