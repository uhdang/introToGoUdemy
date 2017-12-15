package main

import "fmt"

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func main() {
	s := greet("Jane", "Doe")
	fmt.Println(s)
}
