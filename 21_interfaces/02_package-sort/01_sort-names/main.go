package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Before", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("After", studyGroup)
}
