package oop

import "math"

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Shape interface {
	Area() (area float64)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() (area float64) {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() (area float64) {
	return (t.Base * t.Height) * 0.5
}
