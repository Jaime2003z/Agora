package network

type Transport interface {
	Listen(addr string) error
	Dial(addr string) error
	Send(peerID string, data []byte) error
}
