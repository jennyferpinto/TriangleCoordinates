package test

import (
	"testing"

	Tri "github.com/jennyferpinto/triangle"
)

func TestInTriangleCheckForPoint(t *testing.T) {
	p := Tri.NewCoordinate(15, 20)
	triangle := Tri.NewTriangle(Tri.NewCoordinate(22, 8), Tri.NewCoordinate(12, 55), Tri.NewCoordinate(7, 19))
	res := Tri.CheckForPoint(p, triangle)

	if res != true {
		t.Error("Point:", p.FormatPoint(), "is in triangle", triangle.FormattedTriangle())
	}
}

func TestPointOnEdgeInTriangle(t *testing.T) {
	p := Tri.NewCoordinate(7, 19)
	triangle := Tri.NewTriangle(Tri.NewCoordinate(22, 8), Tri.NewCoordinate(12, 55), Tri.NewCoordinate(7, 19))
	res := Tri.CheckForPoint(p, triangle)

	if res != true {
		t.Error("Point:", p.FormatPoint(), "is in triangle", triangle.FormattedTriangle())
	}
}

func TestOutsideTriangleCheckForPoint(t *testing.T) {
	p := Tri.NewCoordinate(1, 7)
	triangle := Tri.NewTriangle(Tri.NewCoordinate(22, 8), Tri.NewCoordinate(12, 55), Tri.NewCoordinate(7, 19))
	res := Tri.CheckForPoint(p, triangle)

	if res != false {
		t.Error("Point:", p.FormatPoint(), "is in triangle", triangle.FormattedTriangle())
	}
}
