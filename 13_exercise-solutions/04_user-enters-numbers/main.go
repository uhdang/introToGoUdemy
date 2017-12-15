// ask a small number + ask a large number
// print the remainder of larger num divided by smaller num

package main

import "fmt"

func solution() {
	var numOne int
	var numTwo int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter a smaller number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, " = ", numOne%numTwo)
}

func main() {

	for isCorrectInput := false; isCorrectInput == false; {
		var smallNum int
		var bigNum int

		fmt.Println("Small Number: ")
		fmt.Scan(&smallNum)
		fmt.Println("Big Number: ")
		fmt.Scan(&bigNum)

		if smallNum < bigNum {
			isCorrectInput = true

			fmt.Printf("Small Number is %v and big number is %v\n", smallNum, bigNum)
			fmt.Printf("Remainder of bigNum divided by smallNum is %v\n", bigNum%smallNum)
		} else {
			fmt.Println("Big number has to be bigger than Small Number")
		}
	}

	solution()
}
