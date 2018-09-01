package geometry

import "testing"

func TestArea(t *testing.T) {

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{4, 5, 3}
		got := triangle.Area()
		expected := 6.00

		if got != expected {
			t.Errorf("got %.2f expected %.2f", got, expected)
		}
	})

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		expected := 72.0

		if got != expected {
			t.Errorf("got %.2f expected %.2f", got, expected)
		}
	})
}

func TestPerimeter(t *testing.T) {
	triangle := Triangle{12.0, 6, 6.0}
	got := Perimeter(triangle)
	expected := 24.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}
