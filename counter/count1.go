package main

import "fmt"

type Count int

func (c *Count) inc() int {
	*c = *c + 1
	return int(*c)
}

func main() {
	var a Count
	fmt.Println(a)
	fmt.Println(a.inc())
	fmt.Println(a.inc())
	fmt.Println(a.inc(), a.inc(), a.inc())
}

/*
$ go run count1.go
0
1
2
3 4 5
*/
