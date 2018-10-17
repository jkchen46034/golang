package main

import "fmt"

type Counter interface {
	Incr() int
}

func main() {
	counter := Count(0)
	fmt.Println(counter)

	onApiHit(100, &counter)

	p("c:", counter) // 100
}

// this will call the Counter n times
// simulating an api hit n times
func onApiHit(n int, c Counter) {
	for i := 0; i < n; i++ {
		c.Incr()
	}
}

// ----

type Count int

// value-receiver

func (c *Count) Incr() int {
	*c = *c + 1
	return int(*c)
}

// helper for printing values
func p(a ...interface{}) {
	fmt.Println(a...)
}
