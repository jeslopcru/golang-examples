package geometry

// GIVEN base and height integers WHEN call Area function THEN result is the area of a triangle
func Area(base float64, height float64) interface{} {
	return (base * height) / 2
}

// GIVEN three side of a triangle  WHEN call Perimeter function THEN result is the perimeter of a triangle
func Perimeter(a float64, b float64, c float64) interface{} {
	return a + b + c
}
