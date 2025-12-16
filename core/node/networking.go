package node

import (
	"bufio"
	"log"

	"github.com/Jaime2003z/Agora/core/ledger"
	"github.com/libp2p/go-libp2p/core/network"
)

const ProtocolID = "/agora/ledger/1.0.0"

func (n *Node) RegisterProtocols() {
	n.Host.SetStreamHandler(ProtocolID, n.handleStream)
}

func (n *Node) handleStream(s network.Stream) {
	log.Println("New stream from:", s.Conn().RemotePeer())

	rw := bufio.NewReadWriter(
		bufio.NewReader(s),
		bufio.NewWriter(s),
	)

	go n.readLoop(rw)
}

func (n *Node) readLoop(rw *bufio.ReadWriter) {
	for {
		msg, err := rw.ReadString('\n')
		if err != nil {
			return
		}
		log.Println("Received:", msg)
	}
}

// BroadcastBlock sends a block to all connected peers
func (n *Node) BroadcastBlock(b *ledger.Block) error {
	// Serialize the block
	data, err := b.Serialize()
	if err != nil {
		return err
	}

	// Add newline for message delimitation
	data = append(data, '\n')

	// Get all connected peers
	peers := n.Host.Network().Peers()

	// Send to each peer
	for _, peerID := range peers {
		// Skip if it's our own ID
		if peerID == n.Host.ID() {
			continue
		}

		// Open a new stream to the peer
		s, err := n.Host.NewStream(n.Ctx, peerID, ProtocolID)
		if err != nil {
			log.Printf("Error opening stream to %s: %v", peerID, err)
			continue
		}
		defer s.Close()

		// Create a buffered writer
		w := bufio.NewWriter(s)

		// Write the block data
		_, err = w.Write(data)
		if err != nil {
			log.Printf("Error writing to stream: %v", err)
			continue
		}

		// Flush the buffer to ensure the data is sent
		err = w.Flush()
		if err != nil {
			log.Printf("Error flushing stream: %v", err)
		}
	}

	return nil
}
