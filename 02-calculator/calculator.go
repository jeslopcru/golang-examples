package calculator

// GIVEN two integers WHEN call Add function THEN result is sum of them
func Add(x int, y int) int {
	return x + y
}

//GIVEN an array WHEN call AddMultiple function THEN result is sum of elements
func AddMultiple(numberList []int) int {
	result := 0
	for _, number := range numberList {
		result = Add(result, number)
	}
	return result
}
