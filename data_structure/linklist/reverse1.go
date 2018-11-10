package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func (list *Node) GenerateList(vals []int) *Node {
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
	list = head
	return list

}

func (list *Node) Push_back(vals ...int) *Node {
	if len(vals) == 0 {
		return list
	}

	var listvals *Node
	l := listvals.GenerateList(vals)

	if list == nil {
		list = l
		return list
	}

	lastNode := list
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = l

	return list
}

func (list *Node) reverse() *Node {
	var prevNode *Node
	var nextNode *Node

	for currentNode := list; currentNode != nil; {
		nextNode = currentNode.Next
		currentNode.Next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}
	list = prevNode
	return list
}

func (list *Node) print() *Node {
	for current := list; current != nil; current = current.Next {
		fmt.Printf("%p %v\n", current, *current)
	}
	fmt.Print("\n")
	return list
}

func main() {
	var list *Node
	list = list.Push_back(1, 2, 3)
	list = list.Push_back(4, 5)

	list.print()

	l := list.reverse()
	l.print()

	l = l.reverse()
	l.print()
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
