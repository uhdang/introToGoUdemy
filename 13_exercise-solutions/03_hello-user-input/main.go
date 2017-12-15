package main

import "fmt"

func solution() {
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}

func main() {

	var input string

	fmt.Print("Give Name: ")

	name, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("My name is %s\n", input)
	fmt.Printf("number of successfully scanned items : %d\n", name)
	solution()
}
