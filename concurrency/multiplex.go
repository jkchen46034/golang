// Goroutines with multiplex pattern, finding prime numbers

package main

import (
	"fmt"
	"math"
	"sync"
)

func multiplex(a chan int, b chan int) chan int {
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
	a := findPrime(2, 40)
	b := findPrime(40, 80)
	c := multiplex(a, b)

	for val := range c {
		fmt.Println(val)
	}
}

/*
$ go run multiplex.go
2
41
3
5
43
47
7
53
59
11
13
17
19
61
23
67
29
71
31
73
37
79
*/
