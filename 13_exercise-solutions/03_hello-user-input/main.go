package main

import "fmt"

func main() {

	var input string

	fmt.Print("Give Name: ")

	name, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("My name is %s\n", input)
	fmt.Printf("number of successfully scanned items : %d\n", name)

}
