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
	list := push_back(nil, 1)
	push_back(list, 2)
	push_back(list, 3)
	push_back(list, 4)
	push_back(list, 5)

	print(list)
	list = reverse(list)
	print(list)
	list = reverse(list)
	print(list)

}

/*
$ go run list.go
0xc00000e1e0 {1 0xc00000e1f0}
0xc00000e1f0 {2 0xc00000e200}
0xc00000e200 {3 0xc00000e210}
0xc00000e210 {4 0xc00000e220}
0xc00000e220 {5 <nil>}

0xc00000e220 {5 0xc00000e210}
0xc00000e210 {4 0xc00000e200}
0xc00000e200 {3 0xc00000e1f0}
0xc00000e1f0 {2 0xc00000e1e0}
0xc00000e1e0 {1 <nil>}

0xc00000e1e0 {1 0xc00000e1f0}
0xc00000e1f0 {2 0xc00000e200}
0xc00000e200 {3 0xc00000e210}
0xc00000e210 {4 0xc00000e220}
0xc00000e220 {5 <nil>}
*/
