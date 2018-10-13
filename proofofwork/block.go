package main

import (
	"time"
)

type Block struct {
	Timestamp uint64
	Data      []byte
	Prehash   []byte
	Hash      []byte
	Nonce     int
}

func NewBlock(data string, Prehash []byte) *Block {
	block := &Block{uint64(time.Now().Unix()), []byte(data), Prehash, []byte{}, 0}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
