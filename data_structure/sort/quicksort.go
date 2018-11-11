package main

import (
	"fmt"
)

func quickSort(a []int, l, r int) {
	if r <= l+1 {
		return
	}
	m := partition(a, l, r)
	quickSort(a, l, m)
	quickSort(a, m, r)
}

func partition(a []int, l, r int) int {
	pivot := r - 1
	right := r - 2
	left := l
	for right >= left {
		for ; right > left; right-- {
			if a[right] < a[pivot] {
				break
			}
		}
		for ; right > left; left++ {
			if a[left] > a[pivot] {
				break
			}
		}
		if right > left {
			a[left], a[right] = a[right], a[left]
		}
		if left == right {
			if a[left] > a[pivot] {
				a[left], a[pivot] = a[pivot], a[left]
			}
			return left + 1
		}
	}
	return -1
}

func main() {
	a := []int{15, 14, 13, 12, 8, 7, 6, 4}
	fmt.Println(a)
	quickSort(a, 0, len(a))
	fmt.Println(a)
}

/*
$ go run quicksort.go
[15 14 13 12 8 7 6 4]
[4 6 7 8 12 13 14 15]
*/
