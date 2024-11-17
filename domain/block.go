package domain

import (
	"fmt"
	"time"
)

type Block struct {
	Nonce        int
	PreviousHash string
	Timestamp    int64
	Transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	block := new(Block)
	block.Timestamp = time.Now().Unix()
	block.Nonce = nonce
	block.PreviousHash = previousHash
	return block
}

func (b *Block) Print() {
	// Print the block
	fmt.Printf("Block: \n")
	fmt.Printf("Nonce:             %d\n", b.Nonce)
	fmt.Printf("Previous Hash:     %s\n", b.PreviousHash)
	fmt.Printf("Timestamp:         %d\n", b.Timestamp)
	fmt.Printf("Transactions:      %s\n", b.Transactions)
}
