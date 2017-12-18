package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "good morining"
	greeting[1] = "bonjour"
	greeting[2] = "buenos dias"
	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[3])
}
