package main

import (
	"fmt"

	"github.com/Jaime2003z/Agora/core/commons"
	"github.com/Jaime2003z/Agora/core/identity"
	"github.com/Jaime2003z/Agora/core/project"
)

func main() {

	// --------------------------------------------------
	// Time Configuration (simulated ticks)
	// --------------------------------------------------
	now := uint64(1000)        // current tick
	futureTick := uint64(2000) // future tick

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
	// CREATE VOTERS
	// --------------------------------------------------
	var voters []*identity.Identity

	for i := 1; i <= 27; i++ {
		voters = append(voters, &identity.Identity{
			PublicKey: fmt.Sprintf("voter_%02d_pubkey", i),
			Location:  commons.LocalityID("CO-CAU"),
		})
	}

	// --------------------------------------------------
	// PHASE 1: VOTES INSIDE THE WINDOW
	// --------------------------------------------------
	fmt.Println("\n=== Phase 1: First 15 votes (inside window) ===")

	for i := 0; i < 15; i++ {
		err := p.TryVote(voters[i], true, 1.0, now)
		if err != nil {
			fmt.Printf("🚫 %s rejected: %v\n", voters[i].PublicKey, err)
		} else {
			fmt.Printf("✅ %s voted YES\n", voters[i].PublicKey)
		}
	}

	// --------------------------------------------------
	// FORCE TIME ADVANCE
	// --------------------------------------------------
	fmt.Println("\n⏩ Forcing time forward (voting window ends)")

	_ = p.EvaluateProjectLifeCycle(futureTick)

	// --------------------------------------------------
	// PHASE 2: LATE VOTES (SHOULD FAIL)
	// --------------------------------------------------
	fmt.Println("\n=== Phase 2: Late voters (should fail) ===")

	for i := 13; i < len(voters); i++ {
		err := p.TryVote(voters[i], true, 1.0, futureTick)
		if err != nil {
			fmt.Printf("🚫 %s rejected: %v\n", voters[i].PublicKey, err)
		} else {
			fmt.Printf("⚠️ %s voted unexpectedly\n", voters[i].PublicKey)
		}
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
