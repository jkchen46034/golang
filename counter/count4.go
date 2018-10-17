package main

import (
	"fmt"
	"time"
)

func Bang(energy int) time.Duration {
	return time.Second * time.Duration(energy)
}

func main() {
	anonymous := func(energy int) time.Duration {
		return Bang(energy)
	}

	fmt.Printf("%T\n%T\n", Bang, anonymous)

	// They have the same types
	// Assign the Bang to the anonymous variable
	// with a type of `func(int) time.Duration`.
	anonymous = Bang
	anonymous(10)
}
