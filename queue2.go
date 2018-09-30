// This function implments a linear queue

package main

import (
	"fmt"
)

type Queue struct {
	arr   [5]int
	front int32
	tail  int32
	size  int32
}

func NewQueue() *Queue {
	return &Queue{[5]int{}, -1, -1, 5}
}

func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

func (q *Queue) IsFull() bool {
	return q.tail == q.size-1
}

func (q *Queue) Push(val int) {

	if q.front == -1 {
		q.front = 0
	}

	if q.IsFull() {
		panic("q.push() overflows")
	}

	q.tail = q.tail + 1
	q.arr[q.tail] = val
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		panic("q.Pop() underflows")
	}

	q.front = q.front + 1

	return q.arr[q.front-1]
}

func main() {
	q := NewQueue()
	fmt.Printf("%+v\n", q)
	fmt.Printf("IsFull? %v, IsEmpty? %v\n", q.IsFull(), q.IsEmpty())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	fmt.Printf("%+v\n", q)
	fmt.Printf("IsFull? %v, IsEmpty? %v\n", q.IsFull(), q.IsEmpty())
	// panic
	//q.Push(6)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Printf("%+v\n", q)
	fmt.Printf("IsFull? %v, IsEmpty? %v\n", q.IsFull(), q.IsEmpty())
	// panic
	//q.Pop()
}
