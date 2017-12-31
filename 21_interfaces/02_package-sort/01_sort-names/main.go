package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	// in order to use "Sort", since Sort method accepts type "Interface", we need to convert type "people" into "Interface" type

	sort.Sort(studyGroup)
	fmt.Println("after: ", studyGroup)

	// using sort.Strings

	//studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	//// studyGroup is type people and type people returns string slice.

	//// looking at Strings function from sort package,
	//// it accepts a []string, and returns Sort(StringSlice(a))
	//fmt.Println("before: ", studyGroup)
	//sort.Strings(studyGroup)
	//fmt.Println("after: ", studyGroup)

}
