// This class implements a binary search tree, with insert, find, delete, and rebalance methods

package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

/*
    10
   /  \
  7   12
 / \ / \
4  9 11 15

*/

func NewNode(val int) *Node {
	return &Node{nil, nil, val}
}

func (node *Node) insert(val int) *Node {
	if node == nil {
		return NewNode(val)
	}
	if val > node.val {
		node.right = node.right.insert(val)
	} else if val < node.val {
		node.left = node.left.insert(val)
	}
	return node
}

func (tree *Node) find(val int) *Node {
	if tree == nil || tree.val == val {
		return tree
	}

	if val > tree.val {
		return tree.right.find(val)
	}
	return tree.left.find(val)
}

// Inorder traversal of a bst is sort
func (node *Node) inorder() {
	if node == nil {
		return
	}
	node.left.inorder()
	node.visit()
	node.right.inorder()
}

// Breadth frist traversal

func (n *Node) bfs() {
	if n == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, n)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		n.visit()
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

func (n *Node) visit() {
	fmt.Println(n.val)
}

func main() {
	vals := []int{10, 7, 12, 4, 9, 11, 15}

	var root *Node

	for _, val := range vals {
		root = root.insert(val)
	}

	fmt.Println("Breadth First Traversal:")
	root.bfs()

	fmt.Println("Inorder Traversal:")
	root.inorder()

	var node *Node
	node = root.find(4)
	fmt.Println("trying to find 4, found: ", node.val)
	node = root.find(9)
	fmt.Println("trying to find 9, found: ", node.val)
	node = root.find(13)
	fmt.Println("trying to find 13, found: ")
	if node == nil {
		fmt.Println(node)
	} else {
		fmt.Println(node.val)
	}

	return
}

/*

$ go run bst.go
Breadth First Traversal:
10
7
12
4
9
11
15
Inorder Traversal:
4
7
9
10
11
12
15
trying to find 4, found:  4
trying to find 9, found:  9
trying to find 13, found:
<nil>

*/
