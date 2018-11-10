// This class implements a binary search tree, with Insert, Find, Delete, and Rebalance methods

package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Val   int
	Right *Tree
}

/*
    20
   /  \
  18  22
 /    / \
4   21 25
 \
  9

*/

func NewTree(val int) *Tree {
	return &Tree{nil, val, nil}
}

func (tree *Tree) Insert(val int) *Tree {
	if tree == nil {
		return NewTree(val)
	}
	if val > tree.Val {
		tree.Right = tree.Right.Insert(val)
	} else if val < tree.Val {
		tree.Left = tree.Left.Insert(val)
	}
	return tree
}

func (tree *Tree) Find(val int) *Tree {
	if tree == nil || tree.Val == val {
		return tree
	}

	if val > tree.Val {
		return tree.Right.Find(val)
	}
	return tree.Left.Find(val)
}

// Find the minium value that is equal to or greater than the
// search val
var visited int

func (t *Tree) MinFind(val int) (*Tree, int) {
	visited += 1
	if t == nil {
		return nil, 9999999
	}
	if val == t.Val {
		return t, 0
	}
	if val > t.Val {
		minNode, childmin := t.Right.MinFind(val)
		localmin := t.Val - val
		if localmin > 0 && localmin < childmin {
			return t, localmin
		} else {
			return minNode, childmin
		}
	}

	minNode, childmin := t.Left.MinFind(val)
	localmin := t.Val - val
	if localmin > 0 && localmin < childmin {
		return t, localmin
	} else {
		return minNode, childmin
	}
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
	fmt.Println(`

        20
       /  \
      18  25
      /   / \
     4   23 28
      \
       9

	`)
	vals := []int{20, 18, 25, 4, 9, 23, 28}

	var tree *Tree

	for _, val := range vals {
		tree = tree.Insert(val)
	}

	fmt.Println("Breadth First Traversal:")
	tree.Bfs()

	fmt.Println("Inorder Traversal:")
	tree.Inorder()

	var t *Tree

	t = tree.Find(4)
	fmt.Println("Finding == 4, found: ", t.Val)

	t = tree.Find(9)
	fmt.Println("Finding == 9, found: ", t.Val)

	t = tree.Find(23)
	fmt.Println("Finding == 23, found: ", t.Val)

	t = tree.Find(13)
	fmt.Println("Finding == 13, found: ", t)

	t = tree.Find(24)
	fmt.Println("Finding == 24, found: ", t)

	visited = 0
	var n *Tree
	n, _ = tree.MinFind(13)
	fmt.Println("Finding >= 13, found: ", n.Val)
	n, _ = tree.MinFind(24)
	fmt.Println("Finding >= 24, found: ", n.Val)
	n, _ = tree.MinFind(9)
	fmt.Println("Finding >= 9, found: ", n.Val)
	n, _ = tree.MinFind(6)
	fmt.Println("Finding >= 6, found: ", n.Val)
	n, _ = tree.MinFind(2)
	fmt.Println("Finding >= 2, found: ", n.Val)
	n, _ = tree.MinFind(30)
	fmt.Println("Finding >= 30, found: ", n)
	return
}

/*
$ go run bst.go


        20
       /  \
      18  25
      /   / \
     4   23 28
      \
       9


Breadth First Traversal:
20
18
25
4
23
28
9
Inorder Traversal:
4
9
18
20
23
25
28
Finding == 4, found:  4
Finding == 9, found:  9
Finding 23, found:  23
Finding 13, found:  <nil>
Finding 24, found:  <nil>
Finding >= 13, found:  18
Finding >= 24, found:  25
Finding >= 9, found:  9
Finding >= 6, found:  9
Finding >= 2, found:  4
Finding >= 30, found:  <nil>

*/
