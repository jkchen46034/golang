// A goroutine with generator pattern, finding prime numbers

package main

import (
	"fmt"
	"math"
)

func findPrime() chan int {
	c := make(chan int)
	go func() {
		for i := 2; ; i++ {
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
	}()
	return c
}

func main() {
	c := findPrime()

	for i := 0; ; i++ {
		val := <-c
		fmt.Println(val)
		if val*val > 10000 {
			break
		}
	}
}

/*
$ go run generator.go
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
41
43
47
53
59
61
67
71
73
79
83
89
97
101
*/
