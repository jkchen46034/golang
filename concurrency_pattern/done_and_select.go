// done

package main

import (
	"fmt"
	"math"
)

func multiplex(a chan int, b chan int, done chan int) chan int {
	c := make(chan int)
	go func() {
		for n := 2; n > 0; {
			select {
			case val := <-a:
				c <- val
			case val := <-b:
				c <- val
			case <-done:
				n--
			}
		}
		close(c)
	}()

	return c
}

func findPrime(from int, to int, done chan int) chan int {
	c := make(chan int)
	go func() {
		for i := from; i < to; i++ {
			if i < 2 {
				continue
			}
			mid := int(math.Sqrt(float64(i)))
			isPrime := true
			for j := 2; j <= mid; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				c <- i
			}
		}
		done <- 1
	}()
	return c
}

func main() {
	done := make(chan int)
	a := findPrime(1, 10, done)
	b := findPrime(10, 20, done)
	c := multiplex(a, b, done)

	for val := range c {
		fmt.Println(val)
	}
}

/*
$ go run done.go
2
3
5
11
7
13
17
19
*/
