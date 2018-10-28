// Mutex, implemented in channel
package main

import (
	"fmt"
	"time"
)

func Adder1(a *int, m chan int) {
	<-m
	*a = *a + 1
	m <- 1
}

func Adder2(a *int, m chan int) {
	<-m
	*a = *a + 2
	m <- 1
}

func main() {
	m := make(chan int)
	a := 0
	for i := 0; i < 1000; i++ {
		go Adder1(&a, m)
		go Adder2(&a, m)
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
