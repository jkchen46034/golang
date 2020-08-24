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

	node0.DFS_Iteration()
}

func (t *Node) DFS_Iteration() {
	if t == nil {
		return
	}
	stack := NewStack()
	stack.Push(t)
	for stack.Len() > 0 {
		n := stack.Pop()
		fmt.Println("visited: ", n.val)
		if n.right != nil {
			stack.Push(n.right)
		}
		if n.left != nil {
			stack.Push(n.left)
		}
	}
}

type Stack []*Node

func NewStack() *Stack {
	s := Stack(make([]*Node, 0))
	return &s
}

func (s *Stack) Push(val ...*Node) {
	*s = append(*s, val...)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Pop() *Node {
	length := len(*s)
	val := (*s)[length-1]
	*s = (*s)[0 : length-1]
	return val
}
