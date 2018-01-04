package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// Why is this deadlock? and how can i fix it?
