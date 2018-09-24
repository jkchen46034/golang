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

func Sort(name []string, score []int, comp func(i, j int) bool) {
	n := len(score)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if comp(i, j) {
				score[i], score[j] = score[j], score[i]
				name[i], name[j] = name[j], name[i]
			}
		}
	}
}

func main() {
	name := []string{"Jack", "Benny", "Alfred", "JK"}
	score := []int{13, 12, 8, 12}

	print(name, score)

	// Sorted by score
	Sort(name, score, func(i, j int) bool {
		if score[j] > score[i] {
			return true
		} else {
			return false
		}
	})
	print(name, score)

	// Sorted by name
	Sort(name, score, func(i, j int) bool {
		if strings.ToLower(name[j]) < strings.ToLower(name[i]) {
			return true
		} else {
			return false
		}
	})
	print(name, score)
}

/*
$ go run player1.go
Jack - 13
Benny - 12
Alfred - 8
JK - 12
Jack - 13
Benny - 12
JK - 12
Alfred - 8
Alfred - 8
Benny - 12
Jack - 13
JK - 12
*/
