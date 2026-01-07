package network

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
)

type Host interface {
	ID() peer.ID
	Stop() error
	Start() error
	Broadcast(msg []byte)
	Send(ctx context.Context, peer peer.ID, msg []byte) error
}
