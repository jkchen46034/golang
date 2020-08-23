package main

import (
	"fmt"
)

func binarySearch(a []int, v int, left, right int) (bool, int) {
	for left <= right {
		mid := (left + right) / 2
		if v == a[mid] {
			return true, mid
		} else if v < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false, -1
}

func main() {
	a := []int{2, 3, 4, 7, 9, 12, 15, 28}
	fmt.Println(a)
	exist, index := binarySearch(a, 4, 0, len(a) - 1)
	fmt.Println("Trying to find 4: ", exist, index)
	exist, index = binarySearch(a, 28, 0, len(a) - 1)
	fmt.Println("Trying to find 28: ", exist, index)
	exist, index = binarySearch(a, 0, 0, len(a) - 1)
	fmt.Println("Trying to find 0: ", exist, index)
}
