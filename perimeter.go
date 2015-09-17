package main

import "math"

//Circle is an exported struct
type Circle struct {
	x float64
	y float64
	r float64
}

//Rect is an exported struct
type Rect struct {
	x float64
	y float64
}

//Shape is an exported struct
type Shape interface {
	perimeter() float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rect) perimeter() float64 {
	return 2 * (r.x + r.y)
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}
