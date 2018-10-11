// This file implements a stack with array which is always fixed size
// does not panic

package main

import (
	"fmt"
)

type Stack struct {
	arr [5]int
	top int // point to the index of topest element
}

func (s *Stack) Push(val int) {
	if s.IsFull() {
		panic("stack overflows")
	}

	s.top = s.top + 1
	s.arr[s.top] = val
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("stack underflows")
	}
	val := s.arr[s.top]
	s.top = s.top - 1
	return val
}

func (s *Stack) IsEmpty() bool {
	return s.top <= -1
}

func (s *Stack) IsFull() bool {
	return s.top >= len(s.arr)-1
}

func main() {
	s := Stack{[5]int{}, -1}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Printf("%v\n", s)
	fmt.Printf("is empty? %v\n", s.IsEmpty())
	fmt.Printf("is full? %v\n", s.IsFull())

	// stack overflows
	// s.Push(5)

	val := s.Pop()
	fmt.Println(val)
	val = s.Pop()
	fmt.Println(val)
	val = s.Pop()
	fmt.Println(val)
	fmt.Printf("%v\n", s)
	fmt.Printf("is empty? %v\n", s.IsEmpty())
	fmt.Printf("is full? %v\n", s.IsFull())
	val = s.Pop()
	fmt.Println(val)
	val = s.Pop()
	fmt.Println(val)
	fmt.Printf("%v\n", s)
	fmt.Printf("is empty? %v\n", s.IsEmpty())
	fmt.Printf("is full? %v\n", s.IsFull())

	// stack underflows
	//s.Pop()
}

/*
$ go run stack2.go
{[1 2 3 4 5] 4}
is empty? false
is full? true
5
4
3
{[1 2 3 4 5] 1}
is empty? false
is full? false
2
1
{[1 2 3 4 5] -1}
is empty? true
is full? false

*/
