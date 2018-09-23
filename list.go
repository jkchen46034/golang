package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func push_back(n *Node, vals ...int) *Node {

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
		return n
	}

	// nil root
	if n == nil {
		return headToNewList
	}

	// concatenate with the existing list pointed to by n
	current = n
	for current.next != nil {
		current = current.next
	}
	current.next = headToNewList

	return n
}

func reverse(n *Node) *Node {
	var prev *Node
	current := n
	var next *Node

	for current != nil {
		// change link
		next = current.next
		current.next = prev

		// move ahead by one link
		prev = current
		current = next
	}
	head := prev
	return head
}

func print(n *Node) {
	for ; n != nil; n = n.next {
		fmt.Printf("%p %v\n", n, *n)
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
