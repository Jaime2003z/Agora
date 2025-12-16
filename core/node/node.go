package node

import (
	"context"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	host "github.com/libp2p/go-libp2p/core/host"
)

type Node struct {
	Host host.Host
	Ctx  context.Context
}

func NewNode() (*Node, error) {
	ctx := context.Background()

	h, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/0",
		),
	)
	if err != nil {
		return nil, err
	}

	log.Println("Node started with ID:", h.ID())

	return &Node{
		Host: h,
		Ctx:  ctx,
	}, nil
}
