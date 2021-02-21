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
	node3 := &Node{3, nil, nil}
	node4 := &Node{4, nil, node7}
	node8 := &Node{8, nil, nil}
	node5 := &Node{5, nil, nil}
	node2 := &Node{2, node8, node5}
	node1 := &Node{1, node3, node4}
	node0 := &Node{0, node1, node2}

	fmt.Println(`	
        0
       / \
      1   2
      /\  / \
     3  4 8  5
        \    
        7    
	`)

	node0.DFS()
}

func (t *Node) DFS() {
	if t == nil {
		return
	}
	t.left.DFS()
	fmt.Println(t.val)
	t.right.DFS()
}
