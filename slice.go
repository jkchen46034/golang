// This file demonstrates how a slice can "lose" memory. Be aware!!!!

package main

import (
	"fmt"
)

func Print(s []int) {
	fmt.Printf("len %d cap %d, slice %v\n", len(s), cap(s), s)
}

func main() {
	s := make([]int, 0)
	Print(s)
	s = append(s, 0)
	Print(s)
	s = append(s, 1)
	Print(s)
	s = append(s, 1)
	Print(s)
	s = append(s, 1)
	Print(s)
	s = s[1:len(s)]
	Print(s)
	s = s[1:len(s)]
	Print(s)
	s = s[1:len(s)]
	Print(s)
	s = s[1:len(s)]
	Print(s)
}

/*

Note: see that capacity of slice decreasing?
That is, space under slice are becoming inaccessiable.

golang$ go run slice.go
len 0 cap 0, slice []
len 1 cap 1, slice [0]
len 2 cap 2, slice [0 1]
len 3 cap 4, slice [0 1 1]
len 4 cap 4, slice [0 1 1 1]
len 3 cap 3, slice [1 1 1]
len 2 cap 2, slice [1 1]
len 1 cap 1, slice [1]
len 0 cap 0, slice []

*/
