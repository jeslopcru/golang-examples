package geometry

import "math"

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (triangle Triangle) Area() interface{} {
	semiperimeter := (triangle.SideA + triangle.SideB + triangle.SideC) / 2
	radicand := semiperimeter * (semiperimeter - triangle.SideA) * (semiperimeter - triangle.SideB) * (semiperimeter - triangle.SideC)
	return math.Sqrt(radicand)
}

type Rectangle struct {
	SideA float64
	SideB float64
}

func (rectangle Rectangle) Area() interface{} {
	return rectangle.SideA * rectangle.SideB
}

// GIVEN three side of SideA triangle  WHEN call Perimeter function THEN result is the perimeter of SideA triangle
func Perimeter(aTriangle Triangle) interface{} {
	return aTriangle.SideA + aTriangle.SideB + aTriangle.SideC
}
