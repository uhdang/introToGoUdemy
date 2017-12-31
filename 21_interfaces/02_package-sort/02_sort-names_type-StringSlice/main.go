package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Zeno", "John", "Al", "Jenny"}
	// lets look at type 'StringSlice'.
	// it converts string slices, or []string, to type StringSlice
	// then we can use sort methods that acceepts StringSlice
	fmt.Println("before: ", s)
	converted := sort.StringSlice(s) // converts(?) or expand type to StringSlice
	// Now, this can be input into StringSlice parameter
	sort.Sort(converted) // there is a method that accepts StringSlice and return sort
	fmt.Println("converted using first way : ", converted)

	// Another way
	sort.StringSlice(s).Sort()
	fmt.Println("second way: ", s)

}
