package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func push_back(head *Node, vals ...int) *Node {
	// construct the link list for storing vals
	var prev *Node
	var headToNewList *Node
	var current *Node
	for _, val := range vals {
		current = &Node{val, nil}
		if prev == nil {
			headToNewList = current
		} else {
			prev.next = current
		}
		prev = current
	}

	// nothing added
	if headToNewList == nil {
		return head
	}

	// nil root
	if head == nil {
		return headToNewList
	}

	// concatenate with the existing list pointed to by n
	current = head
	for current.next != nil {
		current = current.next
	}
	current.next = headToNewList

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
	list := push_back(nil, 1, 2, 3)
	list = push_back(list, 4, 5)

	print(list)

	list = reverse(list)
	print(list)

	list = reverse(list)
	print(list)
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
