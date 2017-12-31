// n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	// ======== #1 using IntSlice

	//sort.IntSlice(n).Sort()
	//fmt.Println("using IntSlice first way: ", n)

	// ======== #2 using IntSlice

	//sort.Sort(sort.IntSlice(n))
	//fmt.Println("#2", n)

	// ======== #3 using sort.Ints

	//fmt.Println("before: ", n)
	//sort.Ints(n)
	//fmt.Println("after: ", n)

}
