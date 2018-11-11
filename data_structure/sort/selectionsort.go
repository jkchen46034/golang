package main

import (
	"fmt"
)

func selectsort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func main() {
	a := []int{15, 7, 2, 12, 28, 9, 3, 8, 4}
	fmt.Println(a)
	selectsort(a)
	fmt.Println(a)
}

/*
$ go run selectionsort.go 
[15 7 2 12 28 9 3 8 4]
[2 3 4 7 8 9 12 15 28]
*/
