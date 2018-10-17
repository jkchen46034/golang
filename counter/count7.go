package main

import (
	"fmt"
	"time"
)

func main() {
	start := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("concurrent func: ends")
	}

	/*
	   This starts a new concurrent func that doesn't block
	   the main func's execution.
	*/
	go start()

	fmt.Println("main: continues...")

	// TODO: remove the sleep call here and see what happens.
	time.Sleep(5 * time.Second)

	fmt.Println("main: ends")
}
