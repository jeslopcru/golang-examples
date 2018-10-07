package more_geometry

import "testing"

func TestArea(t *testing.T) {

	areaDataProvider := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Triangle{4, 5, 3}, 6.0},
	}

	for _, sut := range areaDataProvider {
		got := sut.shape.Area()
		if got != sut.want {
			t.Errorf("got %.2f want %.2f", got, sut.want)
		}
	}
}
