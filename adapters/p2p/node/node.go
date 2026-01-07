package node

import (
	"github.com/Jaime2003z/Agora/adapters/p2p/consensus"
	"github.com/Jaime2003z/Agora/adapters/p2p/mempool"
	"github.com/Jaime2003z/Agora/adapters/p2p/network"
)

type Node struct {
	ID        string
	Host      network.Host
	Config    *Config
	Mempool   *mempool.Mempool
	Consensus *consensus.Engine
}
