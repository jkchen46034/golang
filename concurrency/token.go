// Implement a token pattern, sequencing a series of tasks
package main

import (
	"fmt"
)

func performTask(in chan int, out chan int) {
	val := <-in
	// do something ....
	val = val + 3
	//...
	out <- val
}

func main() {
	begin := make(chan int)
	in := begin
	out := in
	// build the work chain
	for i := 0; i < 100; i++ {
		out = make(chan int)
		go performTask(in, out)
		in = out
	}
	// now trigger the work
	begin <- 0
	val := <-out
	fmt.Println(val)
}

/*
$ go run token.go
300
*/
