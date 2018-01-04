package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2) // edited

	go func() {
		//wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		//wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
