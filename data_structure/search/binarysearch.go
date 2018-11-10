package main

import (
	"fmt"
)

func binarySearch(a []int, v int, r int, l int) bool {
	if l > r {
		mid := (l + r) / 2
		if v == a[mid] {
			return true
		} else if v > a[mid] {
			return binarySearch(a, v, mid+1, l)
		} else {
			return binarySearch(a, v, r, mid)
		}
	}
	return false
}

func main() {
	a := []int{2, 3, 4, 7, 9, 12, 15, 28}
	fmt.Println(a)
	fmt.Println("Trying to find 9,", binarySearch(a, 9, 0, len(a)))
	fmt.Println("Trying to find 28,", binarySearch(a, 28, 0, len(a)))
	fmt.Println("Trying to find 4,", binarySearch(a, 4, 0, len(a)))
	fmt.Println("Trying to find 5,", binarySearch(a, 5, 0, len(a)))
	fmt.Println("Trying to find 13,", binarySearch(a, 13, 0, len(a)))
	fmt.Println("Trying to find 0,", binarySearch(a, 0, 0, len(a)))
	fmt.Println("Trying to find 100,", binarySearch(a, 100, 0, len(a)))
}

/*
$ go run binarysearch.go
[2 3 4 7 9 12 15 28]
Trying to find 9, true
Trying to find 28, true
Trying to find 4, true
Trying to find 5, false
Trying to find 13, false
Trying to find 0, false
Trying to find 100, false
*/
