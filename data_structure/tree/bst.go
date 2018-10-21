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

func (tree *Node) insert(val int) *Node {
	node := NewNode(val)
	if tree == nil {
		return node
	}
	node0 := tree

	// search
	for node0 != nil {
		if val > node0.val {
			if node0.right == nil {
				node0.right = node
				break
			} else {
				node0 = node0.right
			}
		}
		if val < node0.val {
			if node0.left == nil {
				node0.left = node
				break
			} else {
				node0 = node0.left
			}
		}
	}
	return tree
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
        //visit(n)
		fmt.Println(n.val)
        if n.left != nil {
            queue = append(queue, n.left) 
        }
        if n.right != nil {
            queue = append(queue, n.right)
        }
    }
}


func main() {
	vals := []int{10, 7, 12, 4, 9, 11, 15}

	var tree *Node

	for _, val := range vals {
		tree = tree.insert(val)
	}
	bfs(tree)
	return
}
