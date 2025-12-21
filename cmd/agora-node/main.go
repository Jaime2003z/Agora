package main

import (
	"fmt"
	"math/rand"

	"github.com/Jaime2003z/Agora/core/commons"
	"github.com/Jaime2003z/Agora/core/identity"
	"github.com/Jaime2003z/Agora/core/project"
)

func main() {

	// --------------------------------------------------
	// Time Configuration (simulated ticks)
	// --------------------------------------------------
	now := uint64(1000)        // inside voting window
	futureTick := uint64(2000) // after voting window

	// --------------------------------------------------
	// CREATE PROJECT
	// --------------------------------------------------
	p := &project.Project{
		ID:          "proj-vote-001",
		Title:       "Community Water Project",
		Description: "Build a local water system",
		Status:      project.Propose,
		Proposer:    "proposer_pubkey",
		Location:    commons.LocalityID("CO-CAU"),
		VotingWindow: commons.TimeWindow{
			Start: 900,
			End:   1500,
		},
		Votes: []project.Vote{},
	}

	// --------------------------------------------------
	// CREATE VOTERS (mixed localities)
	// --------------------------------------------------
	var voters []*identity.Identity

	localities := []commons.LocalityID{
		"CO-CAU", // only one valid
		"CO-ANT",
	}

	for i := 1; i <= 300; i++ {
		loc := localities[i%len(localities)] // rotate localities

		voters = append(voters, &identity.Identity{
			PublicKey: fmt.Sprintf("voter_%02d_pubkey", i),
			Location:  loc,
		})
	}

	// --------------------------------------------------
	// PHASE 1: VOTES INSIDE WINDOW (mixed validity)
	// --------------------------------------------------
	for i := 0; i < 278; i++ {
		approve := rand.Intn(2) == 1
		p.TryVote(voters[i], approve, 1.0, now)
	}

	// --------------------------------------------------
	// FORCE TIME ADVANCE
	p.EvaluateProjectLifeCycle(futureTick)

	// PHASE 2: LATE VOTES (time invalid)
	for i := 278; i < len(voters); i++ {
		p.TryVote(voters[i], true, 1.0, futureTick)
	}

	// --------------------------------------------------
	// FINAL PROJECT STATE
	// --------------------------------------------------
	fmt.Println("\n=== Final Project State ===")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Project ID     :", p.ID)
	fmt.Println("Title          :", p.Title)
	fmt.Println("Description    :", p.Description)
	fmt.Println("Status         :", p.Status)
	fmt.Println("Location       :", p.Location)
	fmt.Println("Voting Window  :", p.VotingWindow.Start, "→", p.VotingWindow.End)
	fmt.Println("Total Votes    :", len(p.Votes))

	approve := 0
	reject := 0

	for _, v := range p.Votes {
		if v.Approve {
			approve++
		} else {
			reject++
		}
	}

	fmt.Println("Approve Votes  :", approve)
	fmt.Println("Reject Votes   :", reject)
	fmt.Println("--------------------------------------------------")
}
