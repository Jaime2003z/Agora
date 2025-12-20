package main

import (
	"fmt"
	"log"

	"github.com/Jaime2003z/Agora/core/identity"
	"github.com/Jaime2003z/Agora/core/project"
)

func main() {
	voters := []*identity.Identity{}
	for i := 1; i <= 3; i++ {
		voter, err := identity.NewIdentity(
			fmt.Sprintf("voter%d_public_key", i),
			fmt.Sprintf("LOC-%d", i),
		)
		if err != nil {
			log.Fatalf("Error creating voter %d: %v", i, err)
		}
		voters = append(voters, voter)
	}

	location := project.Location{
		Level:        project.Municipal,
		Country:      "CO",
		State:        "ANT",
		Municipality: "Medellín",
	}

	project, err := project.NewProject(
		"proj-vote-001",
		"New Project",
		"This project will have votes",
		voters[0].PublicKey,
		location,
	)
	if err != nil {
		log.Fatalf("Error creating project: %v", err)
	}

	printProjectStatus(project)

	fmt.Println("\n=== Starting voting ===")

	votes := []struct {
		voterIndex int
		approve    bool
		weight     float32
	}{
		{1, true, 1.0},
		{2, false, 1.0},
	}

	for _, vote := range votes {
		voter := voters[vote.voterIndex]
		_, err := project.CastVote(voter.PublicKey, vote.approve, vote.weight)
		if err != nil {
			log.Printf("Error casting vote: %v", err)
			continue
		}
		fmt.Printf("Vote registered: %s - Approved: %v\n",
			voter.PublicKey[:8]+"...", vote.approve)
	}

	fmt.Println("\n=== Voting results ===")
	printProjectStatus(project)
}

func printProjectStatus(p *project.Project) {
	fmt.Println("\nProject status:")
	fmt.Printf("ID: %s\n", p.ID)
	fmt.Printf("Title: %s\n", p.Title)
	fmt.Printf("Status: %s\n", p.Status)
	fmt.Printf("Total votes: %d\n", len(p.Votes))
	fmt.Printf("Voters: %v\n", p.Voters)
}
