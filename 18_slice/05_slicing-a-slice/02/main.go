package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning",
		"Bonjoiur!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2]) // [Bonjour!, dias!]
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2]) // [Good morning, Bonjour!, dias!]
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:]) // [Selamat pagi!, Gutten morgen]
	fmt.Print("[:] ")
	fmt.Println(greeting[:]) // print out all
}
