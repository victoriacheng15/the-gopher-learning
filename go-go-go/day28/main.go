package main

import (
	"fmt"
	"math"
)

// Day 28: Interfaces
// an interface defines a set of methods.
// the difference between interface and struct
// interface: define behavior without implementation
// struct: define a custom data type

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printGeometry(g geometry) {
	fmt.Println("Size:", g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
}

func detectCircle(g geometry) {
	if circleShape, ok := g.(circle); ok {
		fmt.Println("The shape is circle and radius is", circleShape.radius)
	} else {
		fmt.Println("This is not a circle")
	}
}

func main() {
	rect := rectangle{width: 3, height: 4}
	circ := circle{radius: 5}

	fmt.Println("--- Geometry ---")
	printGeometry(rect)
	fmt.Println()
	printGeometry(circ)

	fmt.Println("\n--- Detect Circle ---")
	detectCircle(rect)
	fmt.Println()
	detectCircle(circ)
}
