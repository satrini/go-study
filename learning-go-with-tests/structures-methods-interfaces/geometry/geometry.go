package geometry

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
