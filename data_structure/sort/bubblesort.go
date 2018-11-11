package main

import (
	"fmt"
)

func bubblesort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{15, 14, 13, 12, 8, 7, 6, 4}
	fmt.Println(a)
	bubblesort(a)
	fmt.Println(a)
}

/*
$ go run bubblesort.go
[15 14 13 12 8 7 6 4]
[4 6 7 8 12 13 14 15]
*/
