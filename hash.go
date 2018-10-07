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

func NewSHA256Hash(inputs ...[]byte) []byte {
	h := sha256.New()
	for _, input := range inputs {
		h.Write(input)
	}
	return h.Sum(nil)
}

func main() {
	h := NewSHA256Hash([]byte("Block"), []byte("1"))
	fmt.Printf("%X\n", h)
}
