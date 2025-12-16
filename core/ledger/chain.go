package ledger

import (
	"errors"
	"sort"
)

type Chain struct {
	Blocks []Block
}

func (c *Chain) AddBlock(b *Block) error {
	if !ValidateBlock(b, c.LastBlock()) {
		return errors.New("invalid block")
	}
	c.Blocks = append(c.Blocks, *b)
	return nil
}

func (c *Chain) LastBlock() *Block {
	if len(c.Blocks) == 0 {
		return nil
	}
	return &c.Blocks[len(c.Blocks)-1]
}

func (c *Chain) GetBlockByID(id int) *Block {
	i := sort.Search(len(c.Blocks), func(i int) bool { return c.Blocks[i].ID >= id })
	if i < len(c.Blocks) && c.Blocks[i].ID == id {
		return &c.Blocks[i]
	}
	return nil
}
