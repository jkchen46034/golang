// This file reverses a linked list

package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func generateList(vals []int) *Node {
	var prev *Node
	var head *Node
	for _, val := range vals {
		node := &Node{val, nil}
		if prev == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return head

}

func push_back(head *Node, vals ...int) *Node {
	if len(vals) == 0 {
		return head
	}

	list := generateList(vals)

	if head == nil {
		return list
	}

	lastNode := head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = list

	return head
}

func reverse(head *Node) *Node {
	var prevNode *Node
	var nextNode *Node

	for currentNode := head; currentNode != nil; {
		nextNode = currentNode.Next
		currentNode.Next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}
	return prevNode
}

func print(head *Node) {
	for current := head; current != nil; current = current.Next {
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
$ go run reverse.go
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
