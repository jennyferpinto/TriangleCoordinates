package triangle

import (
	"fmt"
	"strconv"
)

// Input data:
// + coordinates of triangle vertices
// + coordinates of 4th point
// Develop a program which finds out if 4th point is inside (or outside) the defined triangle

// Coordinate type for storing given coordinates
type Coordinate struct {
	x int
	y int
}

// Triangle type
type Triangle struct {
	A Coordinate
	B Coordinate
	C Coordinate
}

// NewCoordinate constructor for Coordinate struct
func NewCoordinate(x int, y int) Coordinate {
	c := Coordinate{x, y}
	return c
}

// NewTriangle constructor for Triangle struct
func NewTriangle(a Coordinate, b Coordinate, c Coordinate) Triangle {
	t := Triangle{a, b, c}
	return t
}

// CheckForPoint takes in a given point and triangle to compute whether the point is inside the triangle or not
func CheckForPoint(p Coordinate, tr Triangle) bool {

	dX := p.x - tr.C.x
	dY := p.y - tr.C.y
	dX21 := tr.C.x - tr.B.x
	dY12 := tr.B.y - tr.C.y
	D := dY12*(tr.A.x-tr.C.x) + dX21*(tr.A.y-tr.C.y)
	s := dY12*dX + dX21*dY
	t := (tr.C.y-tr.A.y)*dX + (tr.A.x-tr.C.x)*dY

	if D < 0 {
		return s <= 0 && t <= 0 && s+t >= D
	}
	return s >= 0 && t >= 0 && s+t <= D

	// // this solution requires more comparisons and recomputation of values
	// // projecting the chosen point onto each triangle segment by using the dot product.
	// firstSide := (p.x-t.B.x)*(t.A.y-t.B.y) - (t.A.x-t.B.x)*(p.y-t.B.y)
	// secondSide := (p.x-t.C.x)*(t.B.y-t.C.y) - (t.B.x-t.C.x)*(p.y-t.C.y)
	// thirdSide := (p.x-t.A.x)*(t.C.y-t.A.y) - (t.C.x-t.A.x)*(p.y-t.A.y)

	// // the computed values need to be all positive or all negative
	// return (firstSide < 0.0) && (secondSide < 0.0) && (thirdSide < 0.0) ||
	// 	(firstSide > 0.0) && (secondSide > 0.0) && (thirdSide > 0.0)
}

func (c *Coordinate) FormatPoint() string {
	coordinateString := "{" + strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + "}"
	return coordinateString
}

func (t *Triangle) FormattedTriangle() string {
	s := t.A.FormatPoint() + t.B.FormatPoint() + t.C.FormatPoint()
	return s
}

func main() {
	point := Coordinate{15, 20}
	t1 := NewTriangle(Coordinate{22, 8}, Coordinate{12, 55}, Coordinate{7, 19})

	if CheckForPoint(point, t1) {
		fmt.Println("INSIDE TRIANGLE")
	} else {
		fmt.Println("OUTSIDE TRIANGLE")
	}
}
