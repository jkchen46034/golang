package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func push_back(n *Node, v int) *Node {
	if n == nil {
		return &Node{v, nil}
	}

	root := n
	for n.next != nil {
		n = n.next
	}

	n.next = &Node{v, nil}

	return root
}

func reverse(n *Node) *Node {
	var nPrev *Node
	for n != nil {
		next := n.next
		n.next = nPrev
		nPrev = n
		n = next
	}
	return nPrev
}

func print(n *Node) {
	for ; n != nil; n = n.next {
		fmt.Print(n.val)
	}
	fmt.Print("\n")
}

func main() {
	list := push_back(nil, 1)
	push_back(list, 2)
	push_back(list, 3)
	push_back(list, 4)
	print(list)
	print(reverse(list))
}

/*
$ go run list.go
1234
4321
*/
