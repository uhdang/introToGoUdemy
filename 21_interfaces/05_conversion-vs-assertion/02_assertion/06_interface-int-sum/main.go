package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val.(int) + 6) // now it works with assertion. assertion only wokrs for interface !!
}
