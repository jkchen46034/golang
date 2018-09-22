package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)
	// push 1 2 3 to stack
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println("pushed 1, 2, 3, to stack, ", s)
	// pop
	var v int
	v = s[len(s)-1]
	s = s[0 : len(s)-1]
	fmt.Println("poped from stack, ", v)
	// pop
	v = s[len(s)-1]
	s = s[0 : len(s)-1]
	fmt.Println("poped from stack, ", v)
	// the remaing queue
	fmt.Println("stack:", s)
}
/*
$ go run stack1.go
pushed 1, 2, 3, to stack,  [1 2 3]
poped from stack,  3
poped from stack,  2
stack: [1]

*/
