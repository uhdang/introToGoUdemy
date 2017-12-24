package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1) // transform into slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T \n", bs) // byte form
	fmt.Println(string(bs))
}
