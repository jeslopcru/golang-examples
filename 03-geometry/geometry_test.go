package geometry

import "testing"

func TestArea(t *testing.T) {
	triangle := Triangle{4, 5, 3}
	got := Area(triangle)
	expected := 6.00

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestPerimeter(t *testing.T) {
	triangle := Triangle{12.0, 6, 6.0}
	got := Perimeter(triangle)
	expected := 24.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}
