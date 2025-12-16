package node

import (
	"log"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const Rendezvous = "agora-mdns"

func (n *Node) SetupDiscovery() error {
	svc := mdns.NewMdnsService(n.Host, Rendezvous, n)
	return svc.Start()
}

func (n *Node) HandlePeerFound(pi peer.AddrInfo) {
	log.Println("Discovered peer:", pi.ID)
	n.Host.Connect(n.Ctx, pi)
}
