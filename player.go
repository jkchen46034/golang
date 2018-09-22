package main

import (
	"fmt"
   "strings"
)

func print(name []string, score []int) {
	for i, _ := range name {
		fmt.Println(name[i], "-", score[i])
	}
}

func SortByScore(name []string, score []int) {
   n := len(score)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if score[j] > score[i] {
				// swap
				scorei := score[i]
				score[i] = score[j]
				score[j] = scorei
				// swap the other array
				namei := name[i]
				name[i] = name[j]
				name[j] = namei
			}
		}
	}
}

func SortByName(name []string, score []int) {
   n := len(score)
	for i := 0; i < n-1; i++ {
		for j := i+1; j<n; j++ {
			if strings.ToLower(name[j]) < strings.ToLower(name[i]) {
				// swap
				scorei := score[i]
				score[i] = score[j]
				score[j] = scorei
				// swap the other array
				namei := name[i]
				name[i] = name[j]
				name[j] = namei
			}
		}
	}
}

func main() {
	name := make([]string, 0)
	name = append(name, "Jack")
	name = append(name, "Benny")
	name = append(name, "Alfred")
	name = append(name, "JK")

	score := make([]int, 0)
	score = append(score, 13)
	score = append(score, 12)
	score = append(score, 8)
	score = append(score, 12)

   fmt.Println("original:")
	print(name, score)

   fmt.Println("\nby score:")
	SortByScore(name, score)
	print(name, score)

   fmt.Println("\nby name:")
	SortByName(name, score)
	print(name, score)
}

/*

$ go run player.go
original:
Jack - 13
Benny - 12
Alfred - 8
JK - 12

by score:
Jack - 13
Benny - 12
JK - 12
Alfred - 8

by name:
Alfred - 8
Benny - 12
Jack - 13
JK - 12

*/
