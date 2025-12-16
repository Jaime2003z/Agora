package ledger

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash helper
func hash(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// MerkleRoot builds a simple Merkle Root from event payloads
func MerkleRoot(events [][]byte) string {
	if len(events) == 0 {
		return ""
	}

	// Hash leaves
	var level []string
	for _, e := range events {
		level = append(level, hash(e))
	}

	// Build tree
	for len(level) > 1 {
		var next []string

		for i := 0; i < len(level); i += 2 {
			if i+1 < len(level) {
				next = append(
					next,
					hash([]byte(level[i]+level[i+1])),
				)
			} else {
				// Duplicate last if odd
				next = append(
					next,
					hash([]byte(level[i]+level[i])),
				)
			}
		}
		level = next
	}

	return level[0]
}
