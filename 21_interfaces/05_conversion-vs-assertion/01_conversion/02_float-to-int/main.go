package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(int(y))
	fmt.Println(int(y) + x) // while original x and y are different type, converting y into 12 has enabled sum of two values
}
