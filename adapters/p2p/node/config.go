package node

type Config struct {
	NodeID         string
	ListenAddress  string
	BootstrapPeers []string
	DataDir        string
	NetworkID      string
}
