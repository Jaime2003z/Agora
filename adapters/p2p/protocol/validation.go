package protocol

import "errors"

// ==============================
// Message validation
// ==============================

func ValidateMessage(msg Message) error {
	if msg.ID == "" {
		return errors.New("protocol: missing message ID")
	}

	if msg.Type == "" {
		return errors.New("protocol: missing message type")
	}

	if !isKnownMessageType(msg.Type) {
		return errors.New("protocol: unknown message type")
	}

	if msg.Timestamp <= 0 {
		return errors.New("protocol: invalid timestamp")
	}

	if msg.SenderID == "" {
		return errors.New("protocol: missing sender ID")
	}

	if len(msg.Payload) == 0 {
		return errors.New("protocol: empty payload")
	}

	return nil
}

func isKnownMessageType(t MessageType) bool {
	switch t {
	case
		MessageProjectProposal,
		MessageVote,
		MessageSyncRequest,
		MessageSyncResponse:
		return true
	default:
		return false
	}
}
