package domain

import "fmt"

type Blockchain struct {
	TransactionPool []string
	Chain           []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	genesisBlock := NewBlock(0, "initHash")
	bc.Chain = []*Block{genesisBlock}
	return bc
}

func (b *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	block := NewBlock(nonce, previousHash)
	b.Chain = append(b.Chain, block)
	return block
}

func (b *Blockchain) Print() {
	for i, block := range b.Chain {
		fmt.Printf("Chain %d: \n", i)
		block.Print()
	}
}
