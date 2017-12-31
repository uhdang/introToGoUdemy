// reverse order - s := []string{"Zeno", "John", "Al", "Jenny"}
package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// need to expand given into type Interface

func main() {
	s := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("before: ", s)
	sort.Sort(sort.Reverse(s))
	fmt.Println("after: ", s)

}
