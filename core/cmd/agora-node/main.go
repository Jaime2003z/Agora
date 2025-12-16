package main

import (
    "fmt"
    "github.com/Jaime2003z/Agora/core/identity"
    "github.com/Jaime2003z/Agora/core/ledger"
    "github.com/Jaime2003z/Agora/core/governance"
    "github.com/Jaime2003z/Agora/core/storage"
)

func main() {
    // Initialize storage
    db := storage.NewDB("mvp.db")
    defer db.Close()

    // Initialize ledger
    chain := ledger.NewChain(db)

    // Create a test user
    user := identity.NewUser("Alice")
    fmt.Printf("User created: %s (ID: %s)\n", user.Name, user.ID)

    // Create a test proposal
    prop := governance.NewProposal("Clean Local Park", "Project to clean the local park")
    fmt.Printf("Proposal created: %s\n", prop.ID)

    // Vote on the proposal
    governance.Vote(prop, user.ID, true)
    fmt.Printf("Vote registered for proposal %s by %s\n", prop.ID, user.ID)

    // Execute proposal if quorum is met
    if governance.CheckQuorum(prop) {
        governance.ExecuteProposal(prop, chain)
        fmt.Printf("Proposal executed: %s\n", prop.ID)
    } else {
        fmt.Println("Not enough votes yet")
    }
}
