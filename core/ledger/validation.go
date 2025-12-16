package ledger

import "time"

// ValidateBlock checks if a new block is valid by verifying:
// 1. The block's previous hash matches the previous block's hash
// 2. The block's timestamp is not too far in the future
// 3. The block's hash is valid
// 4. The block's ID is one more than the previous block's ID
func ValidateBlock(newBlock, prevBlock *Block) bool {
	if prevBlock == nil {
		// Genesis block case
		return newBlock.PrevHash == "" && newBlock.ID == 0
	}

	// Check block ID increments correctly
	if newBlock.ID != prevBlock.ID+1 {
		return false
	}

	// Check previous hash matches
	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}

	// Check timestamp is not too far in the future (within 1 minute)
	blockTime := time.Unix(0, newBlock.Timestamp)
	if blockTime.After(time.Now().Add(1 * time.Minute)) {
		return false
	}

	// Verify the block's hash is correct
	return newBlock.Hash == newBlock.CalculateHash()
}
