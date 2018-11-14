// multiplex patten implemented with select

package main

import (
	"fmt"
	"math"
)

func multiplex(a chan int, b chan int) chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case val, ok := <-a:
				if !ok {
					a = nil
				} else {
					c <- val
				}
			case val, ok := <-b:
				if !ok {
					b = nil
				} else {
					c <- val
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()

	return c
}

func findPrime(from int, to int) chan int {
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
		close(c)
	}()
	return c
}

func main() {
	a := findPrime(1, 10)
	b := findPrime(10, 20)
	c := multiplex(a, b)

	for val := range c {
		fmt.Println(val)
	}
}

/*
$ go run select.go
11
2
3
13
5
7
17
19

*/
