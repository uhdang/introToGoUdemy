package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions.
without closure, for two or more functions to have access to the same varaible,
that variable would need to be pacakge scope

anonaymous function
a function without a name

func expression
assigning a func to a variable
*/
