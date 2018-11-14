// worker via waitgroup and monitoring

package main

import (
	"fmt"
	"math"
	"sync"
)

func findPrime(wg *sync.WaitGroup, from int, to int, c chan int) {
	go func() {
		defer wg.Done()
		for i := from; i < to; i++ {
			if i < 2 {
				continue
			}
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
	}()
}

func monitor(wg *sync.WaitGroup, c chan int) {
	wg.Wait()
	close(c)
}

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan int)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		findPrime(wg, i*10, (i+1)*10, c)
	}
	go monitor(wg, c)

	for val := range c {
		fmt.Println(val)
	}
}

/*
$ go run worker1.go
2
3
5
7
11
13
17
19
31
37
23
29
*/
