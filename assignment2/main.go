package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height, base float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Printf("The area is %v\n", s.area())
}

func main() {
	s := square{sideLength: 10}
	printArea(s)
	t := triangle{base: 10, height: 5}
	printArea(t)
}
