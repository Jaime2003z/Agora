package project

import "errors"

type ProjectStatus string

const (
	Propose ProjectStatus = "propose"
	Accept  ProjectStatus = "active"
	Reject  ProjectStatus = "reject"
	Close   ProjectStatus = "close"
	Cancel  ProjectStatus = "cancel"
)

type Project struct {
	ID          string
	Title       string
	Description string
	Status      ProjectStatus
	Proposer    string
	Voters      []string
	Milestones  []Milestone
	Proposals   []string
	Votes       []string
}

// constructor
func NewProject(
	id string,
	title string,
	description string,
	proposer string,
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

	return &Project{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      Propose,
		Proposer:    proposer,
		Voters:      []string{},
		Milestones:  []Milestone{},
		Proposals:   []string{},
		Votes:       []string{},
	}, nil
}
