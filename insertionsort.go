package main

import (
	"fmt"
)

func insertionsort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		pivot := arr[i]
		k := i - 1
		for ; k >= 0 && arr[k] > pivot; k = k - 1 {
			if arr[k] > pivot {
				arr[k+1] = arr[k]
			}
		}
		arr[k+1] = pivot
	}
}

func main() {
	arr := []int{15, 2, 7, 12, 28, 3, 9, 4}
	fmt.Println(arr)
	insertionsort(arr)
	fmt.Println(arr)
}
/*
$ go run insertionsort.go 
[15 2 7 12 28 3 9 4]
[2 3 4 7 9 12 15 28]

*/
