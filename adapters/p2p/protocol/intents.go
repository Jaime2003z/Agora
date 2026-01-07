package protocol

import "github.com/Jaime2003z/Agora/core/project"

type ProposalIntent struct {
	MessageID string
	Payload   project.Project
	Timestamp int64
	SenderID  string
}
