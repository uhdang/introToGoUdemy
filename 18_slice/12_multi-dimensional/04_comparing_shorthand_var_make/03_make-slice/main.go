package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student = append(student, "Todd")
	fmt.Println("curious to know the length if it increased: ", len(student)) // since it used append, length increased
	fmt.Println("I'm guessing capacity is now 70? - ", cap(student))          // wow this is interesting. Instead of 70, its 72. So..
	fmt.Println(student)
	fmt.Println(students)
}
