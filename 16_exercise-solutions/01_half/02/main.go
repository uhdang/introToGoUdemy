// using float64 instead of int
// Todd's solution

package main

import "fmt"

func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}
