package main

import "penguino/domain"

func main() {
	blockChain := domain.NewBlockChain()
	blockChain.CreateBlock(5, "hash1")
	blockChain.CreateBlock(2, "hash2")
	blockChain.Print()
}
