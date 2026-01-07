package network

type Peer struct {
	ID      string
	Address string
}

type PeerStore interface {
	Add(peer Peer)
	Remove(peerID string)
	List() []Peer
}
