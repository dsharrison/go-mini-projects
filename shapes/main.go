package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	length float64
}

func main() {

	s := square{5}
	t := triangle{10, 5}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return ((t.base * t.height) / 2)
}

func (s square) getArea() float64 {
	return (s.length * s.length)
}
