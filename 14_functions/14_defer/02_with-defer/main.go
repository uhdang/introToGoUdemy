package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func next() {
	fmt.Println("next")
}

func main() {
	defer world()
	defer next()
	hello()
}
