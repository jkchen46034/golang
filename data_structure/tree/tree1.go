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
    0
   / \
  1   2
 / \   \
3   4   5
 \      /
 7     6

*/

func main() {
	tree3 := &Tree{nil, 3, nil}
	tree4 := &Tree{nil, 4, nil}
	tree6 := &Tree{nil, 6, nil}
	tree8 := &Tree{nil, 8, nil}
	tree5 := &Tree{tree6, 5, nil}
	tree2 := &Tree{tree8, 2, tree5}
	tree1 := &Tree{tree3, 1, tree4}
	tree0 := &Tree{tree1, 0, tree2}

	fmt.Println(`
        0
       / \
      1   2
      /\  / \
     3  4 8  5
            /
           6
    `)

	fmt.Println("Height of tree0: ", Height(tree0))
	fmt.Println("Left and Right height difference of tree0: ", Subtract(tree0))
	fmt.Println("BFS: ", Construct(tree0).BFS())
	fmt.Println(tree0.BFS())
	fmt.Println("Same Structure 0 and copy of 0?  ", SameStructure(tree0, Construct(tree0)))
	fmt.Println("Same Structure 0 and 2?  ", SameStructure(tree0, tree2))
	fmt.Println("Same Structure 2 and 2?  ", SameStructure(tree2, tree2))
	fmt.Println("Same Structure 3 and 2?  ", SameStructure(tree3, tree2))
	fmt.Println("Same Structure 3 and 6?  ", SameStructure(tree3, tree6))

}

func Height(t *Tree) int {
	if t == nil {
		return 0
	}
	return Max(Height(t.Left), Height(t.Right)) + 1
}

func Subtract(t *Tree) int {
	left := Height(t.Left)
	right := Height(t.Right)
	return Max(left-right, right-left)
}

func Construct(t *Tree) *Tree {
	if t == nil {
		return nil
	}
	left := Construct(t.Left)
	right := Construct(t.Right)
	return &Tree{left, t.Val, right}
}

func SameStructure(t0, t1 *Tree) bool {
	if t0 == nil && t1 == nil {
		return true
	}
	if (t0 == nil && t1 != nil) || (t0 != nil && t1 == nil) {
		return false
	}
	return SameStructure(t0.Left, t1.Left) && SameStructure(t0.Right, t1.Right)
}

func (t *Tree) String() string {
	list := t.BFS()
	var s string
	for _, val := range list {
		s = s + string(val) + " "
	}
	return s
}

func (t *Tree) BFS() []int {
	var output []int
	if t == nil {
		return output
	}
	queue := make([]*Tree, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		output = append(output, t.Val)
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
	}
	return output
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
