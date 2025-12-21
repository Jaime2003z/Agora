package project

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Jaime2003z/Agora/core/identity"
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

func (p *Project) TryVote(
	voter *identity.Identity,
	approve bool,
	voterWeight float32,
	currentTick uint64,
) error {

	if currentTick < p.VotingWindow.Start || currentTick > p.VotingWindow.End {
		return ErrProjectNotActive
	}

	// 1️⃣ Primero sincroniza el estado del proyecto
	if err := p.EvaluateProjectLifeCycle(currentTick); err != nil {
		return err
	}

	// 2️⃣ Verifica que aún esté en estado Propose
	if p.Status != Propose {
		return ErrProjectNotActive
	}

	// 3️⃣ Verifica permisos (no muta estado)
	if !p.CanVote(voter) {
		return fmt.Errorf("voter not authorized")
	}

	// 4️⃣ Verifica doble voto
	for _, vote := range p.Votes {
		if vote.VoterID == voter.PublicKey {
			return ErrAlreadyVoted
		}
	}

	// 5️⃣ Registra el voto
	p.Votes = append(p.Votes, Vote{
		VoterID: voter.PublicKey,
		Weight:  voterWeight,
		Approve: approve,
	})

	// 6️⃣ Registra participación

	return nil
}

// CanVote checks if an identity is allowed to vote on a project
func (p *Project) CanVote(i *identity.Identity) bool {
	// Check if project is in correct status
	if p.Status != Propose {
		return false
	}

	// Convert locations to strings for comparison
	voterLoc := string(i.Location)
	projectLoc := string(p.Location)

	// Check location-based voting rights
	if projectLoc == "GLOBAL" {
		return true
	}

	// Check if voter is in the same location or a sub-location
	return voterLoc != "" &&
		(strings.HasPrefix(voterLoc, projectLoc) ||
			strings.HasPrefix(projectLoc, voterLoc))
}

func EvaluateVotingResult(p *Project, currentTick uint64) (bool, error) {
	if currentTick < p.VotingWindow.End {
		return false, errors.New("voting window has not ended")
	}

	totalVotes := len(p.Votes)
	if totalVotes == 0 {
		return false, nil
	}

	approveVotes := 0
	for _, v := range p.Votes {
		if v.Approve {
			approveVotes++
		}
	}

	return float64(approveVotes)/float64(totalVotes) > 0.5, nil
}
