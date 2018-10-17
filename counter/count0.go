package main

import "fmt"

type Count int

func (c Count) inc() int {
	c = c + 1
	return int(c)
}

func main() {
	var a Count
	fmt.Println(a)
	fmt.Println(a.inc())
	fmt.Println(a.inc())
	fmt.Println(a.inc(), a.inc(), a.inc())
}

/*
$ go run count0.go
0
1
1
1 1 1

*/
