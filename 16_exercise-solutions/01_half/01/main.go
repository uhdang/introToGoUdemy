// Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let's us know whether or not the argument was even. For example:
// a. half(1) returns (0, false)
// b. half(2) returns (1, true)

package main

import "fmt"

func half(integer int) (int, bool) {
	firstReturn := integer / 2
	secondReturn := false
	if integer%2 == 0 {
		secondReturn = true
	}
	return firstReturn, secondReturn
}

// Todd's solution

func toddHalf(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(5))
	fmt.Println("----- Todd solution -----")
	// Todd's
	h, even := toddHalf(1)
	i, e := toddHalf(5)

	fmt.Println(h, even)
	fmt.Println(i, e)
}
