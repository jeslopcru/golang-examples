package geometry

import "testing"

func TestArea(t *testing.T) {
	got := Area(12.0, 6)
	expected := 36.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(12.0, 6, 6.0)
	expected := 24.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}
