package main

import (
	"fmt"
)

type Stack []int

func NewStack() Stack {
	return make([]int, 0)
}

func (s *Stack) Push(val ...int) {
	*s = append(*s, val...)
}

func (s *Stack) Pop() int {
	length := len(*s)
	v := (*s)[length-1]
	*s = (*s)[0 : length-1]
	return v
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2, 3)
	fmt.Println("pushed 1, 2, 3, to stack, ", stack)
	v := stack.Pop()
	fmt.Println("poped from stack, ", v)
	v = stack.Pop()
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
