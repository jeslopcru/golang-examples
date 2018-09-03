package more_geometry

import "math"

type Triangle struct {
	a float64
	b float64
	c float64
}

func (triangle Triangle) Area() interface{} {
	semiperimeter := (triangle.a + triangle.b + triangle.c) / 2
	radicand := semiperimeter * (semiperimeter - triangle.a) * (semiperimeter - triangle.b) * (semiperimeter - triangle.c)
	return math.Sqrt(radicand)
}

type Rectangle struct {
	a float64
	b float64
}

func (rectangle Rectangle) Area() interface{} {
	return rectangle.a * rectangle.b
}
