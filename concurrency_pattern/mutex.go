// Mutex, implemented in channel
package main

import (
	"fmt"
	"time"
)

type Mutex struct {
	m chan int
}

func NewMutex() *Mutex {
	return &Mutex{make(chan int)}
}

func (mutex *Mutex) Signal() {
	mutex.m <- 1
}

func (mutex *Mutex) Wait() {
	<-mutex.m
}

func Adder(a *int, amount int, mutex *Mutex) {
	mutex.Wait()
	*a = *a + amount
	mutex.Signal()
}

func main() {
	mutex := NewMutex()
	a := 0
	for i := 0; i < 1000; i++ {
		go Adder(&a, 1, mutex)
		go Adder(&a, 2, mutex)
	}
	mutex.Signal()
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
