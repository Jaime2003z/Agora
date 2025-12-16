package main

import (
    "fmt"
    "core/identity"
    "core/ledger"
    "core/governance"
    "core/storage"
)

func main() {
    // Inicializar almacenamiento
    db := storage.NewDB("mvp.db")

    // Inicializar ledger
    chain := ledger.NewChain(db)

    // Crear identidad de prueba
    user := identity.NewUser("Alice")
    fmt.Printf("User created: %s\n", user.ID)

    // Crear propuesta de ejemplo
    prop := governance.NewProposal("Clean Local Park", "Project to clean the local park")
    fmt.Printf("Proposal created: %s\n", prop.ID)

    // Votar con la identidad de prueba
    governance.Vote(prop, user, true)
    fmt.Printf("Vote registered for proposal %s by %s\n", prop.ID, user.ID)

    // Ejecutar propuesta si cumple quorum
    if governance.CheckQuorum(prop) {
        governance.ExecuteProposal(prop, chain)
        fmt.Printf("Proposal executed: %s\n", prop.ID)
    } else {
        fmt.Println("Not enough votes yet")
    }
}
