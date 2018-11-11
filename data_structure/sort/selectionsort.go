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
	a := []int{15, 2, 7, 12, 28, 3, 9, 4, 8}
	fmt.Println(a)
	selectsort(a)
	fmt.Println(a)
}