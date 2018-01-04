/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	counter := 0
	for n := range factorial(loop(100)) {
		counter++
		fmt.Println(counter, n)
	}
}

func loop(n int) chan int {
	outChannel := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			a := rand.Intn(10)
			outChannel <- a
		}
		close(outChannel)
	}()
	return outChannel
}

func factorial(n chan int) chan int {
	out := make(chan int)
	go func() {
		for m := range n {
			total := 1
			for i := m; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
