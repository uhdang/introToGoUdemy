package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val + 6) // won't work due to different type.
}
