// This function implements a simplied block chain

package main

import (
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prehash := chain.Blocks[len(chain.Blocks)-1].Hash
	block := NewBlock(data, prehash)
	chain.Blocks = append(chain.Blocks, block)
}

func (chain *BlockChain) Print() {
	for i, b := range chain.Blocks {
		fmt.Printf("Block %d\nData: %s\nPrevious Hash: %X\nHash: %X Nonce:%d\n", i, b.Data, b.Prehash, b.Hash, b.Nonce)
	}
}

func (chain *BlockChain) Verify() {
	for i, b := range chain.Blocks {
		pow := NewProofOfWork(b)
		fmt.Println("Block", i, "verified:", pow.Verify())
	}
}
