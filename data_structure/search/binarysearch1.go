package main

import (
	"fmt"
)

func binarySearch(a []int, v int) bool {
	left := 0
	right := len(a)
	for left < right {
		mid := (left + right) / 2
		if v == a[mid] {
			return true
		} else if v <= a[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	a := []int{2, 3, 4, 7, 9, 12, 15, 28}
	fmt.Println(a)
	fmt.Println("Trying to find 9", binarySearch(a, 9))
	fmt.Println("Trying to find 28", binarySearch(a, 28))
	fmt.Println("Trying to find 4", binarySearch(a, 4))
	fmt.Println("Trying to find 7", binarySearch(a, 7))
	fmt.Println("Trying to find 5", binarySearch(a, 5))
	fmt.Println("Trying to find 1", binarySearch(a, 1))
	fmt.Println("Trying to find 13", binarySearch(a, 13))
}

/*
$ go run binarysearch1.go
[2 3 4 7 9 12 15 28]
Trying to find 9 true
Trying to find 28 true
Trying to find 4 true
Trying to find 7 true
Trying to find 5 false
Trying to find 1 false
Trying to find 13 false
*/
