package main

import "fmt"

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

func main() {
	fmt.Println(greet("Jane", "Doe "))
}
