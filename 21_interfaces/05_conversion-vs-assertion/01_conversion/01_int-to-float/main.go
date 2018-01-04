package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(float64(x))
	//  fmt.Println(x + y) // x is int and y is float64. type doesn't match so it throws error
	fmt.Println(y + float64(x))
}
