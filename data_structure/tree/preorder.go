package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	node7 := &Node{7, nil, nil}
	node4 := &Node{4, nil, nil}
	node6 := &Node{6, nil, nil}
	node8 := &Node{8, nil, nil}
	node5 := &Node{5, nil, nil}
	node3 := &Node{3, node6, node7}
	node2 := &Node{2, node8, node5}
	node1 := &Node{1, node3, node4}
	node0 := &Node{0, node1, node2}

	fmt.Println(`
        0
       / \
      1   2
      /\  / \
     3  4 8  5
    / \     
    6 7    
       `)

	node0.PreOrder()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Println("visited ", n.val)
	n.left.PreOrder()
	n.right.PreOrder()
}
