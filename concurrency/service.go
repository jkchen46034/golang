// A goroutine with service pattern, finding prime numbers

package main

import (
	"fmt"
	"math"
)

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
	c := findPrime(2, 20)
	d := findPrime(20, 40)

	for val := range c {
		fmt.Println(val)
	}
	for val := range d {
		fmt.Println(val)
	}
}

/*
$ go run service.go
2
3
5
7
11
13
17
19
23
29
31
37
*/
