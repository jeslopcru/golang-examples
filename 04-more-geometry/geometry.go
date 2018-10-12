package more_geometry

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

type Rectangle struct {
	SideA float64
	SideB float64
}

type Circle struct {
	Radius float64
}

func (triangle Triangle) Area() float64 {
	semiperimeter := (triangle.SideA + triangle.SideB + triangle.SideC) / 2
	radicand := semiperimeter * (semiperimeter - triangle.SideA) * (semiperimeter - triangle.SideB) * (semiperimeter - triangle.SideC)
	return math.Sqrt(radicand)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.SideA * rectangle.SideB
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
