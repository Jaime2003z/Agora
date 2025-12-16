// core/node/node.go
package node

import (
	"context"
	"fmt"
	"log"

	"github.com/Jaime2003z/Agora/core/governance"
	"github.com/Jaime2003z/Agora/core/storage"
	libp2p "github.com/libp2p/go-libp2p"
	host "github.com/libp2p/go-libp2p/core/host"
)

type Node struct {
	Host       host.Host
	Ctx        context.Context
	Governance *governance.GovernanceService
	DB         *storage.DB
}

func (n *Node) ID() string {
	return n.Host.ID().String()
}

func NewNode() (*Node, error) {
	// Initialize database
	db, err := storage.NewDB("agora-node.db")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Create libp2p host
	ctx := context.Background()
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/0",
		),
	)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	log.Println("Node started with ID:", h.ID())

	// Initialize governance service
	govService := governance.NewGovernanceService(db)

	return &Node{
		Host:       h,
		Ctx:        ctx,
		Governance: govService,
		DB:         db,
	}, nil
}

func (n *Node) Close() error {
	return n.DB.Close()
}
