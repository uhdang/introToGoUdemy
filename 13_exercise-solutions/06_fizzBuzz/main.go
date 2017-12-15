// Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".

package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		fmt.Print(i)
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(" - FizzBuzz")
		} else if i%5 == 0 {
			fmt.Print(" - Buzz")
		} else if i%3 == 0 {
			fmt.Print(" - Fizz")
		}
		fmt.Printf("\n")
	}
}
