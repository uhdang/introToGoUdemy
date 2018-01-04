package main

import "fmt"

func main() {
	f := factorial(4)
	//fmt.Println("Total:", <-f)
	for n := range f {
		fmt.Println("Total using range:", n)
	}
}

func factorial(n int) chan int {
	c := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post it to discussion
---- Read other answers from discussion
*/
