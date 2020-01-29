package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		height: 12,
		base:   10,
	}

	s := square{
		sideLength: 10,
	}

	printArea(s)
	printArea(t)

}

func (t triangle) getArea() float64 {
	return math.RoundToEven(0.5 * t.height * t.base)
}

func (s square) getArea() float64 {
	return math.RoundToEven(s.sideLength * s.sideLength)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
