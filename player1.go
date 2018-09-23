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

func Sort(name []string, score []int, comp_swap func(i, j int)) {
	n := len(score)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			comp_swap(i, j)
		}
	}
}

func main() {
	name := []string{"Jack", "Benny", "Alfred", "JK"}
	score := []int{13, 12, 8, 12}

	print(name, score)

	// Sorted by score
	Sort(name, score, func(i, j int) {
		if score[j] > score[i] {
			score[i], score[j] = score[j], score[i]
			name[i], name[j] = name[j], name[i]
		}

	})
	print(name, score)

	// Sorted by name
	Sort(name, score, func(i, j int) {
		if strings.ToLower(name[j]) < strings.ToLower(name[i]) {
			score[i], score[j] = score[j], score[i]
			name[i], name[j] = name[j], name[i]
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
