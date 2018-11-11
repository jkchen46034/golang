package main

import (
	"fmt"
)

func main() {
	var a [1000]int
	a[0] = 1
	a[1] = 1
	for n := 2; n < 1000; n++ {
		a[n] = a[n-1] + a[n-2]
	}
	fmt.Println(a[2])
	fmt.Println(a[3])
	fmt.Println(a[4])
	fmt.Println(a[8])
}

/*
$ go run fib.go
2
3
5
34

*/
