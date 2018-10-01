package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	hash  []byte
}

func NewSHA256Hash(input []byte) []byte {
	h := sha256.New()
	h.Write(input)
	return h.Sum(nil)
}

func makeMerkleTree(blocks [][]byte) *Node {
	length := len(blocks)
	node := make([]*Node, length, length)

	for i, block := range blocks {
		hash := NewSHA256Hash(block)
		node[i] = &Node{nil, nil, NewSHA256Hash(block)}
		fmt.Printf("%x\n", hash)
	}

	return nil
}

func main() {
	blocks := make([][]byte, 0)
	blocks = append(blocks, []byte("Block1"))
	blocks = append(blocks, []byte("Block2"))
	blocks = append(blocks, []byte("Block3"))
	blocks = append(blocks, []byte("Block4"))
	blocks = append(blocks, []byte("Block5"))
	tree := makeMerkleTree(blocks)
	fmt.Println(tree)
}

/*
$ go run merkletree.go
Block1
Block2
Block3
Block4
Block5
40e9b17a3391b5f461b2b96a2e5810a885f088346b901c65ebb5cf8cf7361103
61edd5d6b03c20f764aab3bc4291b162ff48958e316603d4c07548a37872e380
f76000270122d79ba262f8298080d37b8645d471d3885999274deba5caa7f704
ca174e0d60e2740ce87f5d4b303104edbaa26eae3792f2dc9711e2ea527d6b3c
543ef0b11f17bef0e05f7335930915b763e0c227e58455b6d2c9759698ef686a
<nil>

*/
