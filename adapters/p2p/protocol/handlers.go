package protocol

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jaime2003z/Agora/core/commons"
	"github.com/Jaime2003z/Agora/core/project"
)

type Handler func(msg Message) error

type Router struct {
	handlers map[MessageType]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[MessageType]Handler),
	}
}

func (r *Router) Register(msgType MessageType, handler Handler) {
	if handler == nil {
		panic("protocol: handler cannot be nil")
	}
	r.handlers[msgType] = handler
}

func (r *Router) Handle(msg Message) error {
	handler, exists := r.handlers[msg.Type]
	if !exists {
		return fmt.Errorf("protocol: no handler for message type %s", msg.Type)
	}
	return handler(msg)
}

// ==============================
// Handle project proposals
// ==============================
type ProjectProposalPayload struct {
	ID           string
	Title        string
	Description  string
	Proposer     string
	Location     commons.LocalityID
	VotingWindow commons.TimeWindow
}

func ProjectProposalHandler(
	mempool ProposalSink,
) Handler {

	return func(msg Message) error {

		var payload ProjectProposalPayload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return errors.New("protocol: invalid project proposal payload")
		}

		if payload.ID == "" {
			return errors.New("protocol: missing project ID")
		}

		if payload.Proposer == "" {
			return errors.New("protocol: missing proposer")
		}

		if payload.VotingWindow.Start >= payload.VotingWindow.End {
			return errors.New("protocol: invalid voting window")
		}

		proj := project.Project{
			ID:           payload.ID,
			Title:        payload.Title,
			Description:  payload.Description,
			Proposer:     payload.Proposer,
			Location:     payload.Location,
			VotingWindow: payload.VotingWindow,
			Status:       project.Propose,
		}

		intent := ProposalIntent{
			MessageID: msg.ID,
			Payload:   proj,
			Timestamp: msg.Timestamp,
			SenderID:  msg.SenderID,
		}

		return mempool.AddProposal(intent)
	}
}
