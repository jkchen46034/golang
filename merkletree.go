// This file implements a merkel tree from an array of byte slice

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

func makeMerkleTree(blocks [][]byte) *Node {
	length := len(blocks)

	if length <= 0 {
		return nil
	}

	node := make([]*Node, length, length)

	for i, block := range blocks {
		node[i] = &Node{nil, nil, NewSHA256Hash(block)}
		fmt.Printf("%p %v %v %X\n", node[i], node[i].left, node[i].right, node[i].hash)
	}

	// repeatively build the tree until a root node is built
	for length := len(node); length != 1; {
		newLen := length / 2
		if length%2 != 0 {
			newLen = newLen + 1
		}
		newNode := make([]*Node, newLen, newLen)
		for i := 0; i < newLen; i++ {
			if 2*i+1 < length {
				left := node[2*i]
				right := node[2*i+1]
				newNode[i] = &Node{left, right, NewSHA256Hash(left.hash, right.hash)}
			} else {
				newNode[i] = &Node{node[2*i], nil, node[2*i].hash}
			}
		}
		node = newNode
		length = len(node)
	}

	return node[0]
}

// breadth frist traversal
func bfs(n *Node) {
	if n == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, n)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		visit(n)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

func visit(n *Node) {
	fmt.Printf("%p %p %p %X\n", n, n.left, n.right, n.hash)
}

func main() {
	blocks := make([][]byte, 0)
	blocks = append(blocks, []byte("Block1"))
	blocks = append(blocks, []byte("Block2"))
	blocks = append(blocks, []byte("Block3"))
	blocks = append(blocks, []byte("Block4"))
	blocks = append(blocks, []byte("Block5"))
	root := makeMerkleTree(blocks)
	bfs(root)
}

/*
$ go run merkletree.go
0xc000066390 0xc000066330 0xc000066360 B99DCFD4F8E724147BF77E8FC121589CCD909C3431C720677413791C66DBF804
0xc000066330 0xc0000662a0 0xc0000662d0 FCB0738B480496B710ACB1CA746BAA8270FE5ACD442855B4BC495BA6E4AAA0A7
0xc000066360 0xc000066300 0x0 543EF0B11F17BEF0E05F7335930915B763E0C227E58455B6D2C9759698EF686A
0xc0000662a0 0xc0000661b0 0xc0000661e0 96C118EC3CD13C548D1A497F0A147225E7F3BB06EDFB056D464313EFEAEEF7DC
0xc0000662d0 0xc000066210 0xc000066240 8223F3D045AEA0E479B18E96C3B86D0EC333F416CA156A250AE44E51E4307DAF
0xc000066300 0xc000066270 0x0 543EF0B11F17BEF0E05F7335930915B763E0C227E58455B6D2C9759698EF686A
0xc0000661b0 0x0 0x0 40E9B17A3391B5F461B2B96A2E5810A885F088346B901C65EBB5CF8CF7361103
0xc0000661e0 0x0 0x0 61EDD5D6B03C20F764AAB3BC4291B162FF48958E316603D4C07548A37872E380
0xc000066210 0x0 0x0 F76000270122D79BA262F8298080D37B8645D471D3885999274DEBA5CAA7F704
0xc000066240 0x0 0x0 CA174E0D60E2740CE87F5D4B303104EDBAA26EAE3792F2DC9711E2EA527D6B3C
0xc000066270 0x0 0x0 543EF0B11F17BEF0E05F7335930915B763E0C227E58455B6D2C9759698EF686A
*/
