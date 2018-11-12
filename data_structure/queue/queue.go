package main

import (
	"fmt"
)

type Queue []int

func NewQueue() *Queue {
	q := Queue(make([]int, 0))
	return &q
}

func (q *Queue) Push(val ...int) {
	*q = append((*q), val...)
}

func (q *Queue) Pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:len(*q)]
	return v
}

func main() {
	q := NewQueue()
	q.Push(1, 2, 3, 4, 5)
	fmt.Println("pushed 1, 2, 3, 4, 5 to queue")
	fmt.Println("Popped", q.Pop(), q.Pop(), q.Pop(), q.Pop(), q.Pop())
}

/*
$ go run queue.go
pushed 1, 2, 3, 4, 5 to queue
Popped 1 2 3 4 5
*/
