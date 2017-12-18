package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "good morning"
	greeting[1] = "bonjour"
	greeting[2] = "buenos dias"
	greeting[3] = "suprabadham" // this will give index our of range error.

	// In order to add element exceeding length, you have to use "append"

	fmt.Println(greeting[2])

}
