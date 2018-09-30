// This function implments a linear queue

package main

import (
	"fmt"
)

type Queue struct {
	arr   [5]int
	front int32
	tail  int32
	size  uint32
}

func NewQueue() *Queue {
	return &Queue{[5]int{}, -1, -1, 5}
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.tail
}

func (q *Queue) IsFull() bool {
	return uint32(q.tail+1) >= q.size
}

func (q *Queue) Push(val int) {

	if q.IsFull() {
		panic(fmt.Sprintf("q.push(%v) overflows %+v \n", val, q))
	}

	q.tail = q.tail + 1
	q.arr[q.tail] = val
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		panic(fmt.Sprintf("q.Pop() underflows %+v\n", q))
	}

	q.front = q.front + 1

	return q.arr[q.front]
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

/*
$ go run queue2.go
&{arr:[0 0 0 0 0] front:-1 tail:-1 size:5}
IsFull? false, IsEmpty? true
&{arr:[1 2 3 4 5] front:-1 tail:4 size:5}
IsFull? true, IsEmpty? false
1
2
3
4
5
&{arr:[1 2 3 4 5] front:4 tail:4 size:5}
IsFull? true, IsEmpty? true
*/
