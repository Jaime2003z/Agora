package ledger

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	ID         int    `json:"id"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prev_hash"`
	Data       string `json:"data"`
	Timestamp  int64  `json:"timestamp"`
	MerkleRoot string `json:"merkle_root"`
}

// Serialize converts a Block to JSON bytes
func (b *Block) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

// DeserializeBlock converts JSON bytes to a Block
func DeserializeBlock(data []byte) (*Block, error) {
	var block Block
	err := json.Unmarshal(data, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// CalculateHash computes the hash of the block
func (b *Block) CalculateHash() string {
	hasher := sha256.New()
	hasher.Write([]byte(string(b.ID) + b.PrevHash + b.Data + string(b.Timestamp)))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NewBlock creates a new block with the given data and previous hash
func NewBlock(data string, prevHash string) *Block {
	block := &Block{
		Data:      data,
		PrevHash:  prevHash,
		Timestamp: time.Now().Unix(),
	}
	blockData, err := block.Serialize()
	if err != nil {
		// handle error empty for now
		blockData = []byte{}
	}
	block.MerkleRoot = MerkleRoot([][]byte{blockData})
	block.Hash = block.CalculateHash()
	return block
}
