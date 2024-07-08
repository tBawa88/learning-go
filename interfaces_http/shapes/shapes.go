package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t square) getArea() float64 {
	return t.side * t.side
}

// This function doesn't care if the shape is of type triangle or square , it just knows that it will have a method called getArea()float64 on it
func printArea(sh shape) {
	fmt.Printf("Area of this shape is %.2f \n", sh.getArea())
}
