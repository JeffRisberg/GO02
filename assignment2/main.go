package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) float64 {
	return s.getArea()
}

func main() {
	square1 := square{
		sideLength: 5,
	}

	triangle1 := triangle{
		base:   6,
		height: 8,
	}

	fmt.Println(printArea(square1))
	fmt.Println(printArea(triangle1))
}
