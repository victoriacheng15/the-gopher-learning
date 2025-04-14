package main

import "fmt"

// Day 28: Interfaces
// an interace defines a set of methods.
// the different between interface and struct
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
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func printGeometry(g geometry) {
	fmt.Println("Size:", g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
}

func detectCircle(g geometry) {
	if _, ok := g.(circle); ok {
		fmt.Println("The shape is circle and radius is", g.(circle).radius)
	} else {
		fmt.Println("This is not a circle")
	}
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	fmt.Println("--- Geometry ---")
	printGeometry(r)
	fmt.Println()
	printGeometry(c)

	fmt.Println("\n--- Detect Circle ---")
	detectCircle(r)
	fmt.Println()
	detectCircle(c)
}
