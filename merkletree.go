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

func makeMerkleTree(input [][]byte) *Node {
	for _, arr := range input {
		fmt.Println(string(arr))
	}
	list := make([][]byte, len(input))
	for i, arr := range input {
		hash := NewSHA256Hash(arr)
		fmt.Printf("%x\n", hash)
		list[i] = hash
	}

	return nil
}

func main() {
	arr := make([][]byte, 0)
	arr = append(arr, []byte("Block1"))
	arr = append(arr, []byte("Block2"))
	arr = append(arr, []byte("Block3"))
	arr = append(arr, []byte("Block4"))
	arr = append(arr, []byte("Block5"))
	tree := makeMerkleTree(arr)
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
