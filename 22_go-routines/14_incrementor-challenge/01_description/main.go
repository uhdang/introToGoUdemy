package main

import (
	"fmt"
)

func main() {
	counter := 0
	for n := range incrementor(gen("1", "2")) {
		counter++
		fmt.Println(n)
	}

	fmt.Println("Final Counter:", counter)
}

func gen(strings ...string) chan string {
	out := make(chan string)
	go func() {
		for _, n := range strings {
			out <- n
		}
		close(out)
	}()
	return out
}

func incrementor(sc chan string) chan string {
	out := make(chan string)
	for n := range sc {
		go func(n string) {
			for i := 0; i < 20; i++ {
				out <- fmt.Sprint("Process: "+n+" printing:", i)
			}
		}(n)
	}
	close(out)
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
