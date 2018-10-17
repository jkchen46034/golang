package main

import (
	"fmt"
)

/*
Declare a new func type which takes an int and returns an int.
*/
type Cruncher func(int) int

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

	/*
	   Define the crunchers as anonymous funcs with function literals.
	*/
	mul := func(n int) int { return n * 2 }
	add := func(n int) int { return n + 100 }
	sub := func(n int) int { return n - 1 }

	newNums := crunch(nums, mul, add, sub)

	fmt.Println(newNums)

	// Or use them inline:
	newNums = crunch(nums,
		func(n int) int {
			return n * 2
		},
		func(n int) int {
			return n + 100
		},
		func(n int) int {
			return n - 1
		})

	fmt.Println(newNums)
}
