package main

import "fmt"

func main() {
	myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}

/*
You can initilzie map with some data using a "map-literal"
i.e.
commits := map[string]int {
  "rsc": 3711,
  "r": 2138,
  "gri": 1908,
  "adg": 912,
}

The same syntax may be used to initialize an empty map, which is functionally identical to using the `make` function:

i.e.

m = map[string]int{}	is same as    m := make(map[string]int)
*/
