package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Left  *Tree
	Val   int
	Right *Tree
}

func main() {
	tree6 := &Tree{nil, 6, nil}
	tree5 := &Tree{tree6, 5, nil}
	tree4 := &Tree{nil, 4, nil}
	tree3 := &Tree{nil, 3, tree5}
	tree2 := &Tree{tree4, 2, nil}
	tree1 := &Tree{tree2, 1, tree3}
	tree0 := &Tree{nil, 0, tree1}

	fmt.Println(`
        0
         \
          1
         / \
        2   3
       /     \ 
      4       5
             /
            6	
    `)

	fmt.Println("Height of tree0: ", tree0.Height())
	fmt.Println("String of tree0: ", tree0)
	fmt.Println("Left and Right height difference of tree0: ", tree0.Subtract())
	fmt.Println("BFS of tree0:", tree0.BFS())
	fmt.Println("BFS of copy of tree0:", tree0.Construct().BFS())
	fmt.Println("Same Structure 0 and copy of 0?  ", SameStructure(tree0, tree0.Construct()))
	fmt.Println("Same Structure 1 and 3?  ", SameStructure(tree1, tree3))
	fmt.Println("Same Structure 2 and 5?  ", SameStructure(tree2, tree5))
	fmt.Println("Same Structure 3 and 2?  ", SameStructure(tree3, tree2))

	var maxD int
	maxD = 0
	tree0.Diemeter(&maxD)
	fmt.Println("Diemter of tree 0:", maxD)
	maxD = 0
	tree3.Diemeter(&maxD)
	fmt.Println("Diemeter of tree 3: ", maxD)
	maxD = 0
	tree1.Diemeter(&maxD)
	fmt.Println("Diemeter of tree 1: ", maxD)
}

func (t *Tree) Height() int {
	if t == nil {
		return 0
	}
	return Max(t.Left.Height(), t.Right.Height()) + 1
}

func (t *Tree) Subtract() int {
	left := t.Left.Height()
	right := t.Right.Height()
	return Max(left-right, right-left)
}

func (t *Tree) Construct() *Tree {
	if t == nil {
		return nil
	}
	left := t.Left.Construct()
	right := t.Right.Construct()
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

func (t *Tree) Diemeter(maxD *int) int {
	if t == nil {
		return 0
	}
	left := t.Left.Diemeter(maxD)
	right := t.Right.Diemeter(maxD)
	*maxD = Max(left+right+1, *maxD)
	return Max(left, right) + 1
}

func (t *Tree) String() string {
	list := t.BFS()
	var s string
	for _, val := range list {
		s = s + strconv.Itoa(val) + " "
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
