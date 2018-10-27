// Goroutines with multiplex pattern, written in a merge() function

package main

import (
	"fmt"
	"math"
	"sync"
)

func merge(a chan int, b chan int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range a {
			c <- val
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range b {
			c <- val
		}
	}()

	go func() {
		defer close(c)
		wg.Wait()
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
		close(c)
	}()
	return c
}

func main() {
	a := findPrime(2, 20)
	b := findPrime(20, 40)

	for val := range merge(a, b) {
		fmt.Println(val)
	}
}

/*
$ go run merge.go
2
23
3
5
29
7
31
11
13
37
17
19

*/
