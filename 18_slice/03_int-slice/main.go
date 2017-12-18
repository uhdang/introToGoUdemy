package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)

	fmt.Println("-----------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) // 0
	fmt.Println(cap(mySlice)) // 3
	fmt.Println("-----------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
}
