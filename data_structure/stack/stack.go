package main

import (
	"fmt"
)

type Stack []int

func NewStack() *Stack {
	s := Stack(make([]int, 0))
	return &s
}

func (s *Stack) Push(val ...int) {
	*s = append(*s, val...)
}

func (s *Stack) Pop() int {
	length := len(*s)
	val := (*s)[length-1]
	*s = (*s)[0 : length-1]
	return val
}

func main() {
	stack := NewStack()
	stack.Push(1, 2, 3, 4, 5)
	fmt.Println("pushed 1, 2, 3, 4, 5")
	fmt.Println("Popped", stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}

/*
$ go run stack.go
pushed 1, 2, 3, 4, 5
Popped 5 4 3 2 1
*/
