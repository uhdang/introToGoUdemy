// s := []string{"Zeno", "John", "Al", "Jenny"}

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Printf("%v\n", s)
}
