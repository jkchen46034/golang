package main

import (
	"fmt"
)

func binarySearch(a []int, v int) bool {
	n := len(a)
	if n == 0 {
		return false
	}
	if n == 1 {
		if a[0] == v {
			return true
		}
		return false
	}

	mid := n / 2

	found := binarySearch(a[0:mid], v)
	if found {
		return true
	}
	return binarySearch(a[mid:n], v)
}

func main() {
	a := []int{2, 3, 4, 7, 9, 12, 15, 28}
	fmt.Println(a)
	fmt.Println("Trying to find 9", binarySearch(a, 9))
	fmt.Println("Trying to find 28", binarySearch(a, 28))
	fmt.Println("Trying to find 4", binarySearch(a, 4))
	fmt.Println("Trying to find 5", binarySearch(a, 5))
}

/*
$ go run binarysearch.go
[2 3 4 7 9 12 15 28]
Trying to find 9 true
Trying to find 28 true
Trying to find 4 true
Trying to find 5 false
*/
