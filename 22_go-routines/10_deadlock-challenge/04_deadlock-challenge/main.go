package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	//for i := 0; i < 10; i++ {
	//fmt.Println(<-c)
	//}

	for n := range c {
		fmt.Println(n)
	}
}

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
//
// while number of input to channel is set to 10, print out from channel is infinite
//
// Why are all goroutines asleep?
//
// same reason. fetching from channel many more than inputting
//
// How can we fix this?
//
// limit fetching to the same number AND CLOSE IT !!
//
