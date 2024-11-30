package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func printArea(shape Shape) {
	area := shape.area()

	if area == math.Floor(area) {
		fmt.Printf("Shape: %.0f\r\n", area)
	} else {
		fmt.Printf("Shape: %.2f\r\n", area)
	}
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 5, height: 10}

	printArea(circle)
	printArea(rectangle)
}
