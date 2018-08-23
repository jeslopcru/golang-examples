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

	otherSum := Add(3, 2)
	otherExpected := 5

	if otherSum != otherExpected {
		t.Errorf("expected '%d' but got '%d'", otherExpected, otherSum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}
