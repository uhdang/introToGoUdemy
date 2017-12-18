package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno",
	}

	fmt.Println(myGreeting)

	//	delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}

/*
check if key exists in map:
You can easily check if a key exists in a go map by using the ok result:

value, ok := myMap[key]

or

if val, ok := myMap[key]; ok {
  // do something
}

*/
