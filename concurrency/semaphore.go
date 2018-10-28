// Producer consumer via semaphore, implemented using buffered channel
package main

import (
	"fmt"
)

func producer(s chan int) {
	for i := 0; i < 10; i++ {
		s <- i // produce
	}
}

func consumer(s chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-s) // consume
	}
}

func main() {
	s := make(chan int, 3)
	go producer(s)
	go consumer(s)
	for {
	}
}

/*
$ go run semaphore.go
0
1
2
3
4
5
6
7
8
9
*/
