// done channel

package main

import (
	"fmt"
)

func runserver() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("server is running")
		}
		c <- 1
	}()

	return c
}

func main() {
	c := runserver()
	<-c
	fmt.Println("exited from main")
}

/*
$ go run done.go
server is running
server is running
server is running
exited from main
*/
