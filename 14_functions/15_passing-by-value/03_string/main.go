package main

import "fmt"

func changeMe(z string) {
	fmt.Println(z)
	z = "Rocky"
	fmt.Println(z)
}

func main() {
	name := "Todd"
	fmt.Println(name)

	changeMe(name)

	fmt.Println(name)
}
