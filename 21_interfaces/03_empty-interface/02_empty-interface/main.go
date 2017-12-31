package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}       // 100
	fmt.Printf("%T\n", s) // float64
	info(s)               //10, 100
}
