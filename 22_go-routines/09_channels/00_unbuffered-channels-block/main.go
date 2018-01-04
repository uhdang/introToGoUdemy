package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // create a channel 'c'

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // take the number i and put it into a channel 'c'
			// When that happens, it blocks until it reaches to the other end.
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // take the channel off
		}
	}()

	time.Sleep(time.Second)
}
