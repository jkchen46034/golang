package main

import (
	"fmt"
)

type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{make([]int, 0)}
}

func (q *Queue) Push(val int) {
	q.queue = append(q.queue, val)
}

func (q *Queue) Pop() int {
	v := q.queue[0]
	q.queue = q.queue[1:len(q.queue)]
	return v
}

func main() {
	q := NewQueue()
	// 3 pushes
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println("push 1, 2, 3, to queue:", *q)
	// 2 pops
	fmt.Println("pop from queue, ", q.Pop())
	fmt.Println("pop from queue, ", q.Pop())
	fmt.Println("queue: ", *q)
}
/*

$ go run queue.go
push 1, 2, 3, to queue: {[1 2 3]}
pop from queue,  1
pop from queue,  2
queue:  {[3]}

*/
