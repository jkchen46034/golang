// Implement a sequence pattern, sequencing a series of tasks
// Task implemented with generation pattern
package main

import (
	"fmt"
)

func performTask(in chan int) chan int {
	out := make(chan int)
	go func() {
		val := <-in
		// do something ....
		val = val + 3
		//...
		out <- val
		return
	}()
	return out
}

func main() {
	begin := make(chan int)
	in := begin
	out := in
	// build the work chain
	for i := 0; i < 10000; i++ {
		out = performTask(in)
		in = out
	}
	// now trigger the work
	begin <- 0
	fmt.Println(<-out)
}

/*
$ go run token.go
30000
*/
