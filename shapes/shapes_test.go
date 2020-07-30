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

func TestSquareArea(t *testing.T) {
	got := Area(Rectangle{
		Width:  10,
		Height: 10,
	})
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
