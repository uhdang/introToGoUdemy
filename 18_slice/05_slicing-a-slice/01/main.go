package main

import "fmt"

func main() {
	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)       // ["a", "b", "c", "g", "m", "z"]
	fmt.Println(mySlice[2:4])  // ["c", "g"]
	fmt.Println(mySlice[2])    // ["c"]
	fmt.Println("myString"[2]) // ASCI of S
}
