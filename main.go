package main

import (
	"fmt"
	"strconv"
)

// Input data:
// + coordinates of triangle vertices
// + coordinates of 4th point
// Develop a program which finds out if 4th point is inside (or outside) the defined triangle

type Coordinate struct {
	x int
	y int
}

type Triangle struct {
	A Coordinate
	B Coordinate
	C Coordinate
}

// NewTriangle is constructor
func NewTriangle(a Coordinate, b Coordinate, c Coordinate) Triangle {
	t := Triangle{a, b, c}
	return t
}

func CheckForPoint(point Coordinate, t Triangle) bool {

	a := []Coordinate{t.A, t.B, t.C}
	a1 := []Coordinate{point, t.A, t.B}
	a2 := []Coordinate{point, t.A, t.C}
	a3 := []Coordinate{point, t.B, t.C}

	if area(a) == area(a1)+area(a2)+area(a3) {
		return true
	}

	return false
}

func (c *Coordinate) FormatPoint() string {
	coordinateString := "{" + strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + "}"
	return coordinateString
}

func (t *Triangle) FormattedTriangle() string {
	s := t.A.FormatPoint() + t.B.FormatPoint() + t.C.FormatPoint()
	return s
}

func area(c []Coordinate) int {

	// Area A = [ x1(y2 – y3) + x2(y3 – y1) + x3(y1-y2)] / 2

	area := c[0].x*(c[1].y-c[2].y) + c[1].x*(c[2].y-c[0].y) + c[2].x*(c[0].y-c[1].y)

	return area
}

func main() {
	point := Coordinate{20, 0}
	t1 := NewTriangle(Coordinate{0, 0}, Coordinate{10, 30}, Coordinate{20, 0})

	if CheckForPoint(point, t1) == true {
		fmt.Println("INSIDE TRIANGLE")
	} else {
		fmt.Println("OUTSIDE TRIANGLE")
	}
}
