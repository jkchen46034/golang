// This function implments an unlimited circular queue

package main

import (
	"fmt"
)

type Queue struct {
	q     []int
	front int
	tail  int
	size  int
}

func NewQueue(initSize int) *Queue {
	if initSize < 1 {
		panic(fmt.Sprintf("NewQueue(%v), Oh please dont give me a value less than one\n", initSize))
	}
	return &Queue{make([]int, initSize, initSize), -1, -1, initSize}
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Length() int {
	return q.tail - q.front
}

func (q *Queue) IsFull() bool {
	return q.Length() == q.size
}

func (q *Queue) Push(val int) {
	if q.IsFull() {
		q.Copy()
	}

	q.tail = q.tail + 1
	q.q[q.tail%q.size] = val
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		q.Copy()
	}

	q.front = q.front + 1

	return q.q[q.front%q.size]
}

func (q *Queue) Copy() {

	// double the size
	oldSize := len(q.q)
	newSize := oldSize * 2
	newq := make([]int, newSize, newSize)

	// copy
	for i := 0; !q.IsEmpty(); i++ {
		newq[i] = q.Pop()
	}

	// reassign slice reference
	q.q = newq
	q.size = (newSize)
	q.front = -1
	q.tail = oldSize - 1
}

func main() {
	q := NewQueue(3)
	fmt.Printf("%+v\n", q)
	fmt.Printf("IsFull? %v, IsEmpty? %v, Length %d,  %+v\n", q.IsFull(), q.IsEmpty(), q.Length(), q)
	q.Push(1)
	q.Push(2)
	fmt.Printf("Pushed 1,2, IsFull? %v, IsEmpty? %v, Length %d,  %+v\n", q.IsFull(), q.IsEmpty(), q.Length(), q)
	q.Push(3)
	q.Push(4)
	fmt.Printf("Pushed 3,4, IsFull? %v, IsEmpty? %v, Length %d,  %+v\n", q.IsFull(), q.IsEmpty(), q.Length(), q)

	fmt.Printf("Poped %d\n", q.Pop())
	fmt.Printf("Poped %d\n", q.Pop())
	fmt.Printf("Poped %d\n", q.Pop())
	fmt.Printf("Poped %d\n", q.Pop())
	fmt.Printf("IsFull? %v, IsEmpty? %v, Length %d,  %+v\n", q.IsFull(), q.IsEmpty(), q.Length(), q)
	q.Push(5)
	q.Push(6)
	fmt.Printf("Pushed 5, 6, IsFull? %v, IsEmpty? %v, Length %d,  %+v\n", q.IsFull(), q.IsEmpty(), q.Length(), q)
}

/*
$ go run queue4.go
&{q:[0 0 0] front:-1 tail:-1 size:3}
IsFull? false, IsEmpty? true, Length 0,  &{q:[0 0 0] front:-1 tail:-1 size:3}
Pushed 1,2, IsFull? false, IsEmpty? false, Length 2,  &{q:[1 2 0] front:-1 tail:1 size:3}
Pushed 3,4, IsFull? false, IsEmpty? false, Length 4,  &{q:[1 2 3 4 0 0] front:-1 tail:3 size:6}
Poped 1
Poped 2
Poped 3
Poped 4
IsFull? false, IsEmpty? true, Length 0,  &{q:[1 2 3 4 0 0] front:3 tail:3 size:6}
Pushed 5, 6, IsFull? false, IsEmpty? false, Length 2,  &{q:[1 2 3 4 5 6] front:3 tail:5 size:6}

*/
