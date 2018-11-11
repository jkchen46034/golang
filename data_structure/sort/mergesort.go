package main

import (
	"fmt"
)

func mergeSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	m := n / 2
	mergeSort(a[0:m])
	mergeSort(a[m:n])
	combine(a, m)
}

func combine(a []int, m int) {
	for i := m; i < len(a); i++ {
		k := i - 1
		val := a[i]
		for ; k >= 0 && a[k] > val; k-- {
			a[k+1] = a[k]
		}
		a[k+1] = val
	}
}

func main() {
	a := []int{15, 2, 7, 12, 28, 3, 9, 4, 8}
	fmt.Println(a)
	mergeSort(a)
	fmt.Println(a)
}

/*
$ go run mergesort.go
[15 2 7 12 28 3 9 4 8]
[2 3 4 7 8 9 12 15 28]
*/
