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
	return NewBlock("Genesis Bllock", []byte{})
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
		fmt.Printf("%+v %+v %+v\n", b.Data, b.Prehash, b.Hash)
	}
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("sent to J 100 million dollars")
	bc.AddBlock("charged J 110 million dollars")
	fmt.Printf("%+v\n", bc)
	print(bc)
}

/*
$ go run blockchain.go
&{Blocks:[0xc000084000 0xc000084050 0xc0000840a0]}
[71 101 110 101 115 105 115 32 66 108 108 111 99 107] [] [0 215 243 145 92 198 227 22 73 210 144 150 189 156 70 117 60 125 15 169 0 237 153 227 140 37 14 95 228 80 35 112]
[115 101 110 116 32 116 111 32 74 32 49 48 48 32 109 105 108 108 105 111 110 32 100 111 108 108 97 114 115] [0 215 243 145 92 198 227 22 73 210 144 150 189 156 70 117 60 125 15 169 0 237 153 227 140 37 14 95 228 80 35 112] [159 247 170 168 243 101 34 38 146 40 81 102 42 222 53 186 200 251 101 136 227 155 174 140 212 241 41 252 99 114 23 48]
[99 104 97 114 103 101 100 32 74 32 49 49 48 32 109 105 108 108 105 111 110 32 100 111 108 108 97 114 115] [159 247 170 168 243 101 34 38 146 40 81 102 42 222 53 186 200 251 101 136 227 155 174 140 212 241 41 252 99 114 23 48] [101 55 183 52 117 229 27 37 228 151 50 28 2 137 61 83 150 244 154 234 228 207 146 6 251 115 171 172 209 203 154 99]

*/
