package main

import "fmt"

type Shape interface {
	area() float64
}

type Square struct {
	length float64
}

type Triangle struct {
	base   float64
	height float64
}

func (s Square) area() float64 {
	return s.length * s.length
}

func (t Triangle) area() float64 {
	return t.base * t.height / 2
}

func main() {
	t := Triangle{base: 10, height: 10}
	s := Square{length: 10}

	printArea(t)
	printArea(s)
}

func printArea(s Shape) {
	fmt.Println("Area: ", s.area())
}
