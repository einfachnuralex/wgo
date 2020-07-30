package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{
		Width:  10.00,
		Height: 10.00,
	})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// table based tests
func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{10, 50}, 250},
	}
	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %g want %g  in %#v", got, test.want, test.shape)
			}
		})
	}
}

func BenchmarkCircle_Area(b *testing.B) {
	circle := Circle{10}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}
