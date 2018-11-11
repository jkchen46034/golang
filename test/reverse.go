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

func reverse(a []int) []int {
	stack := NewStack()
	for i := 0; i < len(a); i++ {
		stack.Push(a[i])
	}
	for i := 0; i < len(a); i++ {
		a[i] = stack.Pop()
	}
	return a
}

func main() {
	a := []int{3, 4, 5, 6}
	fmt.Println(a)
	fmt.Println(reverse(a))
}

/*
$ go run reverse.go
[3 4 5 6]
[6 5 4 3]

*/
