package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c) // receive from c
}

// Why is this deadlock? How can I fix it?

// Channels have to be in concurrency.
// Code right now is in sequential order, where putting element into a channel and pulling it out is NOT concurrent.
// so, making pushing element into channel and pulling out concurrent, it resolves the problem.

// TODD :: when element is being pushed, original code does not have anything to receive it!!

// by default, sends and received block until the other side is ready.
//
