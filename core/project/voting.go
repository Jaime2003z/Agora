package project

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyVoted     = errors.New("user has already voted on this project")
	ErrProjectNotActive = errors.New("project is not in active status for voting")
)

type Vote struct {
	VoterID string
	Weight  float32
	Approve bool
}

func (p *Project) CastVote(voterID string, approve bool, voterWeight float32) (*Project, error) {

	if p.Status != Propose && p.Status != Accept {
		return nil, fmt.Errorf("%w: current status is %s", ErrProjectNotActive, p.Status)
	}

	for _, voter := range p.Voters {
		if voter == voterID {
			return nil, ErrAlreadyVoted
		}
	}

	voteID := fmt.Sprintf("%s_%s", p.ID, voterID)
	p.Votes = append(p.Votes, voteID)
	p.Voters = append(p.Voters, voterID)

	if p.Status == Propose && len(p.Votes) >= 1 {
		p.Status = Accept
	}

	return p, nil
}

type VotingResult struct {
	TotalVotes   int
	ApproveVotes int
	RejectVotes  int
	QuorumMet    bool
}
