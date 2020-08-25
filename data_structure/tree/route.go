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

	node0.Route(4, 7)
	node0.Route(7, 8)
}

func (n *Node) Route(s int, e int) bool {
	q1 := make([]int, 0)
	n.Find(s, &q1)
	Reverse(&q1)
	fmt.Println(q1)
	q2 := make([]int, 0)
	n.Find(e, &q2)
	Reverse(&q2)
	fmt.Println(q2)
	q := MakeQ(q1, q2)
	fmt.Println("The route from ", s, " to ", e, " is ", q)
	return true
}

func (n *Node) Find(val int, q *[]int) bool {
	if n == nil {
		return false
	}

	fmt.Println("visited  ", n.val)
	if n.val == val {
		fmt.Println("found leaf ", n.val)
		*q = append(*q, n.val)
		return true
	}
	found := n.left.Find(val, q)
	if found == true {
		fmt.Println("node on path ", n.val)
		*q = append(*q, n.val)
		return true
	}
	found = n.right.Find(val, q)
	if found {
		fmt.Println("node on path ", n.val)
		*q = append(*q, n.val)
	}
	return found
}

func Reverse(q *[]int) {
	length := len(*q)
	for i := 0; i < length/2; i++ {
		(*q)[i], (*q)[length-1-i] = (*q)[length-1-i], (*q)[i]
	}
}

func MakeQ(q1 []int, q2 []int) (q []int) {
	i := 0
	for q1[i] == q2[i] {
		i = i + 1
	}
	i = i - 1
	if i < 0 {
		return
	}

	q1r := q1[i:len(q1)]
	q2r := q2[i+1 : len(q2)]
	Reverse(&q1r)
	q = append(q1r, q2r...)
	return q
}
