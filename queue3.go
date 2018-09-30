// This function implments a fixed-size circular queue in golang

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
	return q.front == q.tail
}

func (q *Queue) IsFull() bool {
	return q.tail > q.front && (q.tail-q.front)%q.size == 0
}

func (q *Queue) Push(val int) {

	if q.IsFull() {
		panic(fmt.Sprintf("q.push(%v) overflows %+v \n", val, q))
	}

	q.tail = q.tail + 1
	q.arr[q.tail%q.size] = val
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		panic(fmt.Sprintf("q.Pop() underflows %+v\n", q))
	}

	q.front = q.front + 1

	return q.arr[q.front%q.size]
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
	fmt.Printf("%+v\n", q)
	fmt.Printf("IsFull? %v, IsEmpty? %v\n", q.IsFull(), q.IsEmpty())
	q.Push(6)
	// panic
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Printf("IsFull? %v, IsEmpty? %v %+v\n", q.IsFull(), q.IsEmpty(), q)
	q.Push(7)
	q.Push(8)
	fmt.Println(q.Pop())
	fmt.Printf("IsFull? %v, IsEmpty? %v %+v\n", q.IsFull(), q.IsEmpty(), q)
	fmt.Println(q.Pop())
	fmt.Printf("IsFull? %v, IsEmpty? %v %+v\n", q.IsFull(), q.IsEmpty(), q)
}

/*
$ go run queue3.go
&{arr:[0 0 0 0 0] front:-1 tail:-1 size:5}
IsFull? false, IsEmpty? true
&{arr:[1 2 3 4 5] front:-1 tail:4 size:5}
IsFull? true, IsEmpty? false
1
&{arr:[1 2 3 4 5] front:0 tail:4 size:5}
IsFull? false, IsEmpty? false
2
3
4
5
6
IsFull? false, IsEmpty? true &{arr:[6 2 3 4 5] front:5 tail:5 size:5}
7
IsFull? false, IsEmpty? false &{arr:[6 7 8 4 5] front:6 tail:7 size:5}
8
IsFull? false, IsEmpty? true &{arr:[6 7 8 4 5] front:7 tail:7 size:5}
*/
