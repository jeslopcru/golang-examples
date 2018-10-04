package more_geometry

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

type Rectangle struct {
	a float64
	b float64
}

func (triangle Triangle) Area() float64 {
	semiperimeter := (triangle.a + triangle.b + triangle.c) / 2
	radicand := semiperimeter * (semiperimeter - triangle.a) * (semiperimeter - triangle.b) * (semiperimeter - triangle.c)
	return math.Sqrt(radicand)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.a * rectangle.b
}
