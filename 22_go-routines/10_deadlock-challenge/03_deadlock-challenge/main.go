package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

// Why does this only print zero?

// Because while for loop keeps on adding elements into the channel, there is only one side that receives it.

// What can you do to get it to print all 0 - 9 ?

// create same number of loops that each receives what's been added.
