package main

import (
	"fmt"
)

func binarySearch(a []int, v int, left, right int) bool {
	for left <= right {
		mid := (left + right) / 2
		if v == a[mid] {
			return true
		} else if v < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	a := []int{2, 3, 4, 7, 9, 12, 15, 28}
	fmt.Println(a)
	fmt.Println("Trying to find 4", binarySearch(a, 4, 0, len(a) - 1))
	fmt.Println("Trying to find 15", binarySearch(a, 15, 0, len(a) - 1))
	fmt.Println("Trying to find 28", binarySearch(a, 28, 0, len(a) - 1))
	fmt.Println("Trying to find 2", binarySearch(a, 2, 0, len(a) - 1))
	fmt.Println("Trying to find 8", binarySearch(a, 8, 0, len(a) - 1))
	fmt.Println("Trying to find 30", binarySearch(a, 30, 0, len(a) - 1))
}

/*
$ go run binarysearch.go
[2 3 4 7 9 12 15 28]
Trying to find 4 true
Trying to find 9 true
Trying to find 28 true
Trying to find 8 false
Trying to find 17 false
*/
