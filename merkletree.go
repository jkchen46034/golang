package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	hash  []byte
}

func makeMerkleTree(input [][]byte) *Node {
	for _, arr := range input {
		fmt.Println(string(arr))
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
