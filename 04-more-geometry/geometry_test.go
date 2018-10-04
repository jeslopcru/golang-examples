package more_geometry

import "testing"

func TestArea(t *testing.T) {

	assertArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{4, 5, 3}
		assertArea(t, triangle, 6.00)
	})
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		assertArea(t, rectangle, 72.00)
	})
}
