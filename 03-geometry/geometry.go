package geometry

import "math"

type Triangle struct {
	a float64
	b float64
	c float64
}

// GIVEN three side of a triangle  WHEN call Perimeter function THEN result is the perimeter of a triangle
func Perimeter(aTriangle Triangle) interface{} {
	return aTriangle.a + aTriangle.b + aTriangle.c
}

// GIVEN base and height integers WHEN call Area function THEN result is the area of a triangle
func Area(aTriangle Triangle) interface{} {
	semiperimeter := (aTriangle.a + aTriangle.b + aTriangle.c) / 2
	radicand := semiperimeter * (semiperimeter - aTriangle.a) * (semiperimeter - aTriangle.b) * (semiperimeter - aTriangle.c)
	return math.Sqrt(radicand)
}
