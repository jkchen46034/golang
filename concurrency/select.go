// Goroutines with multiplex pattern, implemented with select statement

package main

import (
	"fmt"
	"math"
	_ "time"
)

var quit chan int

func multiplex(a chan int, b chan int) chan int {
	c := make(chan int)
	go func() {
		cnt := 0
		for {
			select {
			case val := <-a:
				c <- val
			case val := <-b:
				c <- val
			case val := <-quit:
				cnt += val
				if cnt == 2 {
					close(c)
					return
				}
			}
		}
	}()

	return c
}

func findPrime(from int, to int) chan int {
	c := make(chan int)
	go func() {
		for i := from; i < to; i++ {
			isPrime := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				c <- i
			}
		}
		quit <- 1
	}()
	return c
}

func main() {
	quit = make(chan int)
	a := findPrime(2, 10)
	b := findPrime(10, 20)
	c := multiplex(a, b)

	for val := range c {
		fmt.Println(val)
	}
}

/*
$ go run select.go
2
3
5
11
13
7
17
19

*/
