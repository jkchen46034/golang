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
	node9 := &Node{9, nil, nil}
	node7 := &Node{7, node9, nil}
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
         /
	9

	`)

	fmt.Println("tree diemeter is ", node0.right.Diemeter()+node0.left.Diemeter()+1)
}

func (t *Node) Diemeter() int {
	if t == nil {
		return 0
	}
	return Max(t.left.Diemeter(), t.right.Diemeter()) + 1
}

func Max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
