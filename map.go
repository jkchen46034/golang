// This file demonstrates how to structure a hash map for fast retrieval
// while, maintaining a linear strucutre for orderly sequential access

package main

import (
	"fmt"
)

type Golfer struct {
	Fingerprint string
	name        string
	val         []float64
}

type Node struct {
	golfer *Golfer
	index  int
}

func main() {
	// three golfers, two with same fingerprints.
	golfer0 := &Golfer{"342482AD3F", "John Doe", []float64{48.0, 21.3, 120.5}}
	golfer1 := &Golfer{"01234568ABCDE", "Donald Gate", []float64{0.0, 0.0, 1.0, 4.0}}
	golfer2 := &Golfer{"342482AD3F", "Jacky Kim", []float64{27.2, 28.8, 250.4}}

	// golfers stored in a linear list
	node0 := &Node{golfer0, 0}
	node1 := &Node{golfer1, 1}
	node2 := &Node{golfer2, 2}
	slice := []*Node{node0, node1, node2}

	// golfers also stored in a map; for fast retrieval by fingerprint
	gmap := make(map[string][]*Node)
	gmap["342482AD3F"] = append(gmap["342482AD3F"], node0, node2)
	gmap["01234568ABCDE"] = append(gmap["01234568ABCDE"], node1)

	// print our linear list
	fmt.Println("Golfers list:")
	for _, node := range slice {
		fmt.Printf("Golfer %d, %v\n", node.index, *node.golfer)
	}

	// retrieve by fingerprint
	searchKeys := []string{"342482AD3F", "01234568ABCDE", "834A83BD"}
	for _, searchKey := range searchKeys {
		nodes := gmap[searchKey]
		fmt.Printf("List of golfers with fingerprint %s, numbers: %d\n", searchKey, len(nodes))
		for _, node := range nodes {
			fmt.Printf("%v, %v\n", node.index, *node.golfer)
		}
	}
}

/*
$ go run map.go
Golfers list:
Golfer 0, {342482AD3F John Doe [48 21.3 120.5]}
Golfer 1, {01234568ABCDE Donald Gate [0 0 1 4]}
Golfer 2, {342482AD3F Jacky Kim [27.2 28.8 250.4]}
List of golfers with fingerprint 342482AD3F, numbers: 2
0, {342482AD3F John Doe [48 21.3 120.5]}
2, {342482AD3F Jacky Kim [27.2 28.8 250.4]}
List of golfers with fingerprint 01234568ABCDE, numbers: 1
1, {01234568ABCDE Donald Gate [0 0 1 4]}
List of golfers with fingerprint 834A83BD, numbers: 0

*/
