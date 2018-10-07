package more_geometry

import "testing"

func TestArea(t *testing.T) {

	areaDataProvider := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Triangle", Triangle{4, 5, 3}, 6.0},
		{"Circle", Circle{10}, 314.1592653589793},
	}

	for _, sut := range areaDataProvider {
		t.Run(sut.name, func(t *testing.T) {
			got := sut.shape.Area()
			if got != sut.want {
				t.Errorf("got %.2f want %.2f", got, sut.want)
			}
		})
	}
}
