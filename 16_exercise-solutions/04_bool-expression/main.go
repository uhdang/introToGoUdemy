package main

import "fmt"

func main() {
	value := (true && false) || (false && true) || !(false && false)
	fmt.Println("value: ", value)
}
