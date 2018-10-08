// This function implements a simplied block chain

package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

type Block struct {
	Timestamp uint64
	Data      []byte
	Prehash   []byte
	Hash      []byte
}

type BlockChain struct {
	Blocks []*Block
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
	block := &Block{uint64(time.Now().Unix()), []byte(data), Prehash, []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prehash := chain.Blocks[len(chain.Blocks)-1].Hash
	block := NewBlock(data, prehash)
	chain.Blocks = append(chain.Blocks, block)
}

func print(chain *BlockChain) {
	for _, b := range chain.Blocks {
		fmt.Printf("Data: %s\nPrevious Hash: %X\nHash: %X\n", b.Data, b.Prehash, b.Hash)
	}
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("sent to J 100 million dollars")
	bc.AddBlock("charged J 110 million dollars")
	print(bc)
}

/*
$ go run blockchain.go
Data: Genesis Block
Previous Hash:
Hash: 4432A4F78996CA511EB7F7ADD35D61D066558AD0D539B44213A37E266D571A67
Data: sent to J 100 million dollars
Previous Hash: 4432A4F78996CA511EB7F7ADD35D61D066558AD0D539B44213A37E266D571A67
Hash: 7DA2128852033C5D7EEA445022CE9EF9274750C889C2DDFBB1B41B8ECB46FA13
Data: charged J 110 million dollars
Previous Hash: 7DA2128852033C5D7EEA445022CE9EF9274750C889C2DDFBB1B41B8ECB46FA13
Hash: B6BB6E0110CD9F61CF1684922EAF3390F429974EFE29B586FE7B3F5FAD486A69
*/
