package main

import (
	"fmt"
)

func bubblesort(arr []int) {
	n := len(arr)
	var swapped bool
	for i := 0; i < n; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func main() {
	arr := []int{15, 2, 7, 12, 28, 3, 9, 4}
	fmt.Println(arr)
	bubblesort(arr)
	fmt.Println(arr)
}
/*
$ go run bubblesort.go 
[15 2 7 12 28 3 9 4]
[2 3 4 7 9 12 15 28]
*/
