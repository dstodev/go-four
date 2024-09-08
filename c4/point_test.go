package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestPointDestructure(t *testing.T) {
	point := c4.Point{1, 2}

	r, c := point.Get()

	assertEqual(t, 1, r)
	assertEqual(t, 2, c)
}

func TestPointStep(t *testing.T) {
	point := c4.Point{0, 0}

	n := point.Step(c4.North)
	ne := point.Step(c4.NorthEast)
	e := point.Step(c4.East)
	se := point.Step(c4.SouthEast)
	s := point.Step(c4.South)
	sw := point.Step(c4.SouthWest)
	w := point.Step(c4.West)
	nw := point.Step(c4.NorthWest)

	assertEqual(t, c4.Point{-1, 0}, n)
	assertEqual(t, c4.Point{-1, 1}, ne)
	assertEqual(t, c4.Point{0, 1}, e)
	assertEqual(t, c4.Point{1, 1}, se)
	assertEqual(t, c4.Point{1, 0}, s)
	assertEqual(t, c4.Point{1, -1}, sw)
	assertEqual(t, c4.Point{0, -1}, w)
	assertEqual(t, c4.Point{-1, -1}, nw)
}
