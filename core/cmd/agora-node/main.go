package main

import (
    "fmt"
    "core/identity"
    "core/ledger"
    "core/governance"
    "core/storage"
)

func main() {
    db := storage.NewDB("mvp.db")

    chain := ledger.NewChain(db)

    user := identity.NewUser("Alice")
    fmt.Printf("User created: %s\n", user.ID)

    prop := governance.NewProposal("Clean Local Park", "Project to clean the local park")
    fmt.Printf("Proposal created: %s\n", prop.ID)

    governance.Vote(prop, user, true)
    fmt.Printf("Vote registered for proposal %s by %s\n", prop.ID, user.ID)

    if governance.CheckQuorum(prop) {
        governance.ExecuteProposal(prop, chain)
        fmt.Printf("Proposal executed: %s\n", prop.ID)
    } else {
        fmt.Println("Not enough votes yet")
    }
}
