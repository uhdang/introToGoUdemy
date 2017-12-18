package main

import "fmt"

func findGreatest(nums ...int) int {
	theGreatest := 0
	for _, num := range nums {
		if theGreatest < num {
			theGreatest = num
		}
	}
	return theGreatest
}

// Todd Solution

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(findGreatest(1, 2, 3, 4, 5))
	fmt.Println(findGreatest(1, 4, 2, 9, 3, 4))
	// Todd Solution
	fmt.Println("==================")
	greatest := max(4, 9, 123, 543, 432, 98, 32)
	fmt.Println(greatest)

}
