package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Time"] = "Good morining"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}

/*
var m map[string]int

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initilized map. A nil map behaves like an empty map when reading, but attempts to write to anil map will cause a rutime panic; don't do that. To initialize a map, do as it has been done above - use built-in 'make' function

var m = make(map[string]string)

The make function allocates and initializes a hash map data structure and returns a map value that points to it. The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself.
*/
