package main

import (
	"fmt"
)

func mergeSort(a []int) []int{
	n := len(a)
	if n <= 1 {
		return a
	}
	mid := int(float32(n)/2 + 0.5)
	b := a[0:mid]
	c := a[mid:n]

	mergeSort(b)
	mergeSort(c)
	merge(b, c)
	return a
}

func merge(a []int, b []int) {
	m := len(b)
	n := len(a)
	for i := 0; i < m; i++ {
		target := b[i]
		var j int
		for j = 0; j < n && b[i] > a[j]; j++ {
		}

		a = append(a, 0)
		n = n + 1
		copy(a[j+1:n], a[j:n-1])
		a[j] = target
	}
}

func main() {
	a := []int{15, 2, 7, 12, 28, 3, 9, 4, 8}
	fmt.Println(a)
	fmt.Println(mergeSort(a))
}

/*
$ go run mergesort.go
[15 2 7 12 28 3 9 4 8]
[2 3 4 7 8 9 12 15 28]
*/
