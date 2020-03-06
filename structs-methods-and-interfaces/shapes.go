package structs_methods_and_interfaces

import "math"

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) perimeter() (p float64) {
	return 2 * (r.width + r.height)
}

func (r Rectangle) area() (a float64) {
	return r.height * r.width
}

func (c Circle) circumference() (p float64) {
	return
}

func (c Circle) area() (a float64) {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t Triangle) area() (a float64) {
	return (.5 * t.base) * t.height
}
