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
		for j := i + 1; j < n; j++ {
			if score[j] > score[i] {
				score[i], score[j] = score[j], score[i]
				name[i], name[j] = name[j], name[i]
			}
		}
	}
}

func SortByName(name []string, score []int) {
	n := len(score)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if strings.ToLower(name[j]) < strings.ToLower(name[i]) {
				score[i], score[j] = score[j], score[i]
				name[i], name[j] = name[j], name[i]
			}
		}
	}
}

func main() {
	name := []string{"Jack", "Benny", "Alfred", "JK"}
	score := []int{13, 12, 8, 12}

	fmt.Println("Original:")
	print(name, score)

	fmt.Println("\nBy score:")
	SortByScore(name, score)
	print(name, score)

	fmt.Println("\nBy name:")
	SortByName(name, score)
	print(name, score)
}

/*
$ go run player1.go
Original:
Jack - 13
Benny - 12
Alfred - 8
JK - 12

By score:
Jack - 13
Benny - 12
JK - 12
Alfred - 8

By name:
Alfred - 8
Benny - 12
Jack - 13
JK - 12
*/
