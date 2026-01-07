package node

import (
	"time"
)

// ==============================
// Node Lifecycle
// ==============================

// Start initializes the node components
func (n *Node) Start() error {
	// luego: iniciar red, cargar estado, etc
	return nil
}

func (n *Node) Stop() error {
	return nil
}

func (n *Node) Run() {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		n.Tick()
	}
}

func (n *Node) Tick() {
	if n.Consensus != nil {
		n.Consensus.Tick()
	}
}
