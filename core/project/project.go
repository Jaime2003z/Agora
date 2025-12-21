package project

import (
	"errors"
	"strings"

	"github.com/Jaime2003z/Agora/core/commons"
)

type ProjectStatus string

const (
	Propose ProjectStatus = "propose"
	Accept  ProjectStatus = "active"
	Reject  ProjectStatus = "reject"
	Close   ProjectStatus = "close"
	Cancel  ProjectStatus = "cancel"
)

// LocationLevel represents the scope level of a project
//
//go:generate stringer -type=LocationLevel
type LocationLevel int

const (
	International LocationLevel = iota // Project with global scope
	National                           // Project for a specific country
	Regional                           // Project for a specific state/department
	Municipal                          // Project for a specific municipality
)

type Project struct {
	ID           string
	Title        string
	Description  string
	Status       ProjectStatus
	Proposer     string
	Location     commons.LocalityID
	VotingWindow commons.TimeWindow
	Milestones   []Milestone
	Proposals    []string
	Votes        []Vote
}

func isValidLocalityID(l commons.LocalityID) bool {
	if l == "GLOBAL" {
		return true
	}

	parts := strings.Split(string(l), "-")

	// País
	if len(parts[0]) != 2 {
		return false
	}

	// Departamento
	if len(parts) >= 2 && len(parts[1]) < 2 {
		return false
	}

	// Municipio
	if len(parts) >= 3 && len(parts[2]) < 3 {
		return false
	}

	return true
}

// NewProject creates a new project with the given parameters
func NewProject(
	id string,
	title string,
	description string,
	proposer string,
	location *commons.LocalityID,
	currentTick uint64,
	requestedDuration uint64,
) (*Project, error) {

	if id == "" {
		return nil, errors.New("project id is required")
	}

	if title == "" {
		return nil, errors.New("project title is required")
	}

	if proposer == "" {
		return nil, errors.New("proposer is required")
	}

	if !isValidLocalityID(*location) {
		return nil, errors.New("project location is invalid")
	}

	duration := requestedDuration
	if duration < commons.MinTimeWindow {
		duration = commons.MinTimeWindow
	}

	return &Project{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      Propose,
		Proposer:    proposer,
		Location:    *location,
		VotingWindow: commons.TimeWindow{
			Start: currentTick,
			End:   currentTick + duration,
		},
		Milestones: []Milestone{},
		Proposals:  []string{},
		Votes:      []Vote{},
	}, nil
}

func (p *Project) EvaluateProjectLifeCycle(currentTick uint64) error {
	if p.Status != Propose {
		return nil
	}

	if currentTick < p.VotingWindow.End {
		return nil
	}

	approved, err := EvaluateVotingResult(p, currentTick)
	if err != nil {
		return err
	}

	if approved {
		p.Status = Accept
	} else {
		p.Status = Reject
	}

	return nil
}
