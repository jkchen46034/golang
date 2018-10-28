// Mutex, implemented in channel
package main

import (
	"fmt"
	"time"
)

func Adder(a *int, amount int, m chan int) {
	<-m
	*a = *a + amount
	m <- 1
}

func main() {
	m := make(chan int)
	a := 0
	for i := 0; i < 1000; i++ {
		go Adder(&a, 1, m)
		go Adder(&a, 2, m)
	}
	m <- 1
	time.Sleep(10 * time.Millisecond)
	fmt.Println(a)
}

/*
$ go build mutex.go
$ time ./mutex
3000

real	0m0.029s
user	0m0.011s
sys	0m0.017s

*/
