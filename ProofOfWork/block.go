// This function implements a simplied block chain

package main

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Block struct {
	Timestamp uint64
	Data      []byte
	Prehash   []byte
	Hash      []byte
	Nonce     int
}

func (block *Block) SetHash() {
	h := sha256.New()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(block.Timestamp))
	h.Write(b)
	h.Write(block.Data)
	h.Write(block.Prehash)
	block.Hash = h.Sum(nil)
}

func NewBlock(data string, Prehash []byte) *Block {
	block := &Block{uint64(time.Now().Unix()), []byte(data), Prehash, []byte{}, 0}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	//block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
