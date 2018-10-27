// Implement a token pattern, sequencing a series of tasks
// task implemented with generation pattern
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
	for i := 0; i < 100; i++ {
		out = performTask(in)
		in = out
	}
	// now trigger the work
	begin <- 0
	fmt.Println(<-out)
}

/*
$ go run token.go
300
*/
