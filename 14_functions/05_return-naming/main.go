package main

import "fmt"

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

// IMPORTANT - avoid using named returns.
