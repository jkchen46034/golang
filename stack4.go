// Implements a thread safe stack
package main

import (
	"fmt"
)

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{make([]int, 0)}
}

func (s *Stack) Push(val ...int) {
	s.stack = append(s.stack, val...)
}

func (s *Stack) Pop() int {
	length := len(s.stack)
	val := s.stack[length-1]
	s.stack = s.stack[0 : length-1]
	return val
}

func main() {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2, 3)
	fmt.Println("pushed 1, 2, 3, to stack, ", *stack)
	// pop
	fmt.Println("poped from stack, ", stack.Pop())
	// pop
	fmt.Println("poped from stack, ", stack.Pop())
	// the remaing stack
	fmt.Println("stack:", *stack)
}

/*
$ go run stack.go
pushed 1, 2, 3, to stack,  {[1 2 3]}
poped from stack,  3
poped from stack,  2
stack: {[1]}

*/
