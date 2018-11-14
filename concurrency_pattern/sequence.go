// sequence pattern
package main

import (
	"fmt"
)

func performTask(in chan int) chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			val = val + 3
			out <- val
		}
		// If in has been closed, we leave
		close(out)
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
	for i := 0; i < 3; i++ {
		begin <- i
		fmt.Println(<-out)
	}
}

/*

$ go build sequence.go
$ time ./sequence
30000
30001
30002

real	0m0.093s
user	0m0.122s
sys	0m0.045s

*/
