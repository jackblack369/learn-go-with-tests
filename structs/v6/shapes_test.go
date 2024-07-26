package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		desc  string
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, desc: "A", want: 72.0},
		{shape: Circle{Radius: 10}, desc: "B", want: 314.1592653589793},
		{Rectangle{Width: 10, Height: 10}, "C", 100.0},
		{Circle{20}, "D", 1256.6370614359173},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
