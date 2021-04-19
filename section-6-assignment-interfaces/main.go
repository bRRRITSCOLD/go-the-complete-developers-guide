package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2.0
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func main() {
	t := triangle{base: 10, height: 88}
	printArea(t)

	s := square{side: 27}
	printArea(s)
}
