// This class implements a binary search tree, with Insert, Find, delete, and rebalance methods

package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Val   int
}

/*
    10
   /  \
  7   12
 / \ / \
4  9 11 15

*/

func NewTree(Val int) *Tree {
	return &Tree{nil, nil, Val}
}

func (tree *Tree) Insert(Val int) *Tree {
	if tree == nil {
		return NewTree(Val)
	}
	if Val > tree.Val {
		tree.Right = tree.Right.Insert(Val)
	} else if Val < tree.Val {
		tree.Left = tree.Left.Insert(Val)
	}
	return tree
}

func (tree *Tree) Find(Val int) *Tree {
	if tree == nil || tree.Val == Val {
		return tree
	}

	if Val > tree.Val {
		return tree.Right.Find(Val)
	}
	return tree.Left.Find(Val)
}

func (tree *Tree) Inorder() {
	if tree == nil {
		return
	}
	tree.Left.Inorder()
	tree.visit()
	tree.Right.Inorder()
}

func (t *Tree) Bfs() {
	if t == nil {
		return
	}
	queue := make([]*Tree, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		t.visit()
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
	}
}

func (t *Tree) visit() {
	fmt.Println(t.Val)
}

func main() {
	Vals := []int{10, 7, 12, 4, 9, 11, 15}

	var tree *Tree

	for _, Val := range Vals {
		tree = tree.Insert(Val)
	}

	fmt.Println("Breadth First Traversal:")
	tree.Bfs()

	fmt.Println("Inorder Traversal:")
	tree.Inorder()

	var t *Tree

	t = tree.Find(4)
	fmt.Println("trying to Find 4, found: ", t.Val)

	t = tree.Find(9)
	fmt.Println("trying to Find 9, found: ", t.Val)

	t = tree.Find(13)
	fmt.Print("trying to Find 13, found: ")
	if t == nil {
		fmt.Println(t)
	} else {
		fmt.Println(t.Val)
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
trying to Find 4, found:  4
trying to Find 9, found:  9
trying to Find 13, found: <nil>

*/
