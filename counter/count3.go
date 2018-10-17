package main

import (
	"fmt"
)

/*
Declare a new func type which takes an int and returns an int.
*/
type Cruncher func(int) int

/*
Declare funcs as a Cruncher.
*/
func mul(n int) int { return n * 2 }
func add(n int) int { return n + 100 }
func sub(n int) int { return n - 1 }

/*
Crunch func process the numbers and returns the new numbers.

Takes a sequence of numbers in `nums` and variadic number of cruncher funcs.
*/
func crunch(nums []int, a ...Cruncher) (rnums []int) {
	rnums = append(rnums, nums...)

	for _, f := range a {
		for i, n := range rnums {
			rnums[i] = f(n)
		}
	}

	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	newNums := crunch(nums, mul, add, sub)

	fmt.Println(newNums)
}
