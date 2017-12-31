// reverse order - s := []string{"Zeno", "John", "Al", "Jenny"}

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("before: ", s)
	sort.reverse(s).Reverse()
	fmt.Println("after: ", s)

}
