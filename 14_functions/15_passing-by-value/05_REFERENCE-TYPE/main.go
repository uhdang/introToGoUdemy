package main

import "fmt"

func changeMe(z map[string]int) {
	z["Todd"] = 44
}

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Todd"])
}

/*
Allocation with make
Back to allocation. The built-in function make(T, args)
serves a purpose different from new(T). It creates slices, maps, and channels only,
and it returns an initialized (not zeroed) value of type T (not *T). The reason for
the distinction is that these three types represent, under the cover, references to data structures
that must be initialized before use. A slice, for example, is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity, and until those items are initialized, the slice is nil. For slices, maps, and channels, make initialized the internal data structure and prepares the value for use
*/
