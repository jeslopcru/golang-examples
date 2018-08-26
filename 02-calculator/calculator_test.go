package calculator

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}

func TestAddMultiple(t *testing.T) {

	t.Run("array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := AddMultiple(numbers)
		expected := 15

		if expected != got {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})
}

func ExampleAddMultiple() {
	numberList := []int{1, 2, 3}
	sum := AddMultiple(numberList)
	fmt.Println(sum)
	// Output: 6
}
