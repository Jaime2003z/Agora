package network

type Discovery interface {
	Start() error
	Stop() error
	Peers() []Peer
}
