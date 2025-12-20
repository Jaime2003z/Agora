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

type Location struct {
	Level        LocationLevel // Scope level of the project
	Country      string        // ISO 3166-1 alpha-2 country code (e.g., "CO", "PE", "MX")
	State        string        // State/Department code (e.g., "ANT" for Antioquia, "CDMX" for Mexico City)
	Municipality string        // Municipality/City name (e.g., "Medellín", "Lima", "Acapulco")
}

type Project struct {
	ID          string
	Title       string
	Description string
	Status      ProjectStatus
	Proposer    string
	Location    *Location
	Voters      []string
	Milestones  []Milestone
	Proposals   []string
	Votes       []string
}

// NewProject creates a new project with the given parameters
func NewProject(
	id string,
	title string,
	description string,
	proposer string,
	location Location,
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

	if location.Level == 0 {
		return nil, errors.New("project level is required")
	}

	return &Project{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      Propose,
		Proposer:    proposer,
		Location:    &location,
		Voters:      []string{},
		Milestones:  []Milestone{},
		Proposals:   []string{},
		Votes:       []string{},
	}, nil
}
