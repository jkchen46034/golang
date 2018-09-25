// This file demonstrates how to structure a hash map for fast retrieval
// while, maintaining a linear strucutre for orderly sequential access

package main

import (
	"fmt"
)

type Person struct {
	fingerprint string
	name        string
	val         []float64
}

type Node struct {
	person *Person
	index  int
}

func main() {
	// three persons, two with same fingerprints.
	person0 := &Person{"342482AD3F", "John Doe", []float64{48.0, 21.3, 120.5}}
	person1 := &Person{"01234568ABCDE", "Donald Gate", []float64{0.0, 0.0, 1.0, 4.0}}
	person2 := &Person{"342482AD3F", "Jacky Kim", []float64{27.2, 28.8, 250.4}}

	// persons stored in a linear list
	node0 := &Node{person0, 0}
	node1 := &Node{person1, 1}
	node2 := &Node{person2, 2}
	slice := []*Node{node0, node1, node2}

	// persons also stored in a map; for fast retrieval by fingerprint
	pmap := make(map[string][]*Node)
	pmap["342482AD3F"] = append(pmap["342482AD3F"], node0, node2)
	pmap["01234568ABCDE"] = append(pmap["01234568ABCDE"], node1)

	// print our linear list
	fmt.Println("Persons list:")
	for _, node := range slice {
		fmt.Printf("Person %d, %v\n", node.index, *node.person)
	}

	// retrieve by fingerprint
	searchKeys := []string{"342482AD3F", "01234568ABCDE", "834A83BD"}
	for _, searchKey := range searchKeys {
		nodes := pmap[searchKey]
		fmt.Printf("List of persons with fingerprint %s, numbers: %d\n", searchKey, len(nodes))
		for _, node := range nodes {
			fmt.Printf("%v, %v\n", node.index, *node.person)
		}
	}
}

/*
$ go run map.go
Persons list:
Person 0, {342482AD3F John Doe [48 21.3 120.5]}
Person 1, {01234568ABCDE Donald Gate [0 0 1 4]}
Person 2, {342482AD3F Jacky Kim [27.2 28.8 250.4]}
List of persons with fingerprint 342482AD3F, numbers: 2
0, {342482AD3F John Doe [48 21.3 120.5]}
2, {342482AD3F Jacky Kim [27.2 28.8 250.4]}
List of persons with fingerprint 01234568ABCDE, numbers: 1
1, {01234568ABCDE Donald Gate [0 0 1 4]}
List of persons with fingerprint 834A83BD, numbers: 0
*/
