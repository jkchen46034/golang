package main

import (
	"fmt"
)

func main() {
	q := make([]int, 0)
	// 3 pushes
	q = append(q, 1)
	q = append(q, 2)
	q = append(q, 3)
	fmt.Println("push 1, 2, 3, to queue:", q)
	// 2 pops
	var v int
	v = q[0]
	q = q[1:len(q)]
	fmt.Println("pop from queue, ", v)
	v = q[0]
	q = q[1:len(q)]
	fmt.Println("pop from queue, ", v)
	fmt.Println("queue: ", q)
}

/*
$ go run queue1.go
push 1, 2, 3, to queue: [1 2 3]
pop from queue,  1
pop from queue,  2
queue:  [3]

*/
