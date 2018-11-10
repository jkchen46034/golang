package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func push_back(head *Node, vals ...int) *Node {
	if len(vals) == 0 {
		return head
	}

	// construct a link list storing vals
	var prev *Node
	var valList *Node
	var current *Node
	for _, val := range vals {
		current = &Node{val, nil}
		if prev == nil {
			valList = current
		} else {
			prev.next = current
		}
		prev = current
	}

	if head == nil {
		return valList
	}

	// concatenate with head list
	current = head
	for current.next != nil {
		current = current.next
	}
	current.next = valList

	return head
}

func reverse(head *Node) *Node {
	var prev *Node
	current := head
	var next *Node

	for current != nil {
		// change link
		next = current.next
		current.next = prev

		// move ahead by one link
		prev = current
		current = next
	}
	return prev // now head
}

func print(head *Node) {
	for current := head; current != nil; current = current.next {
		fmt.Printf("%p %v\n", current, *current)
	}
	fmt.Print("\n")
}

func main() {
	head := push_back(nil, 1, 2, 3)
	head = push_back(head, 4, 5)

	print(head)

	head = reverse(head)
	print(head)

	head = reverse(head)
	print(head)
}

/*
$ go run list.go
0xc00006e030 {1 0xc00006e040}
0xc00006e040 {2 0xc00006e050}
0xc00006e050 {3 0xc00006e060}
0xc00006e060 {4 0xc00006e070}
0xc00006e070 {5 <nil>}

0xc00006e070 {5 0xc00006e060}
0xc00006e060 {4 0xc00006e050}
0xc00006e050 {3 0xc00006e040}
0xc00006e040 {2 0xc00006e030}
0xc00006e030 {1 <nil>}

0xc00006e030 {1 0xc00006e040}
0xc00006e040 {2 0xc00006e050}
0xc00006e050 {3 0xc00006e060}
0xc00006e060 {4 0xc00006e070}
0xc00006e070 {5 <nil>}

*/
