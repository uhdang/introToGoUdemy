package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "good morning"
	greeting[1] = "bonjour"
	greeting[2] = "buenos dias"

	greeting = append(greeting, "suprahadham")
	greeting = append(greeting, "Zao an")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])   // gidday
	fmt.Println(len(greeting)) // 7
	fmt.Println(cap(greeting)) // 10
}
