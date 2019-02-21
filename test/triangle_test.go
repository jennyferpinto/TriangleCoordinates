package main

import (
	"testing"
)

func TestInTriangleCheckForPoint(t *testing.T) {
	p := Coordinate{20, 0}
	triangle := NewTriangle(Coordinate{0, 0}, Coordinate{10, 30}, Coordinate{20, 0})
	res := CheckForPoint(p, triangle)

	if res != true {
		t.Error("Point:", p.FormatPoint(), "is in triangle", triangle.FormattedTriangle())
	}
}

func TestOutsideTriangleCheckForPoint(t *testing.T) {
	p := Coordinate{0, 20}
	triangle := NewTriangle(Coordinate{0, 0}, Coordinate{10, 30}, Coordinate{20, 0})
	res := CheckForPoint(p, triangle)

	if res != false {
		t.Error("Point:", p.FormatPoint(), "is in triangle", triangle.FormattedTriangle())
	}
}
