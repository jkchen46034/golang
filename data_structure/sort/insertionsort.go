package main

import (
	"fmt"
)

func insertionsort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		pivot := a[i]
		k := i - 1
		for ; k >= 0 && a[k] > pivot; k = k - 1 {
			a[k+1] = a[k]
		}
		a[k+1] = pivot
	}
}

func main() {
	a := []int{15, 2, 7, 12, 28, 3, 9, 4}
	fmt.Println(a)
	insertionsort(a)
	fmt.Println(a)
}

/*
$ go run insertionsort.go
[15 2 7 12 28 3 9 4]
[2 3 4 7 9 12 15 28]

*/
