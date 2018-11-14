// Multiplex pattern

package main

import (
	"fmt"
	"math"
	"sync"
)

func multiplex(a []chan int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		ch := a[i]
		go func() {
			defer wg.Done()
			for val := range ch {
				c <- val
			}
		}()
	}

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
			mid := int(math.Sqrt(float64(i)))
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
	a0 := findPrime(2, 10)
	a1 := findPrime(10, 20)
	a2 := findPrime(20, 30)
	a3 := findPrime(30, 40)

	c := multiplex([]chan int{a0, a1, a2, a3})

	for val := range c {
		fmt.Println(val)
	}
}

/*
$ go run multiplex.go
2
11
23
31
37
3
13
29
5
7
17
19
*/
