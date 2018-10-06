package main

import (
	"fmt"
)

type Stack []int

func NewStack() Stack {
	return make([]int, 0)
}

func (s Stack) Push(val ...int) Stack {
	return append(s, val...)
}

func (s Stack) Pop() (int, Stack) {
	length := len(s)
	return s[length-1], s[0 : length-1]
}

func main() {
	stack := NewStack()

	stack = stack.Push(1)
	stack = stack.Push(2, 3)
	fmt.Println("pushed 1, 2, 3, to stack, ", stack)
	var v int
	v, stack = stack.Pop()
	fmt.Println("poped from stack, ", v)
	v, stack = stack.Pop()
	fmt.Println("poped from stack, ", v)
	fmt.Println("stack:", stack)
}

/*
$ go run stack.go
pushed 1, 2, 3, to stack,  {[1 2 3]}
poped from stack,  3
poped from stack,  2
stack: {[1]}

*/
