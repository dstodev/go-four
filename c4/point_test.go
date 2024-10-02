package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestNewPoint(t *testing.T) {
	point := c4.NewPoint(1, 2)

	util.AssertEqual(t, 1, point.Row)
	util.AssertEqual(t, 2, point.Column)
}

func TestPointDestructure(t *testing.T) {
	point := c4.Point{1, 2}

	r, c := point.Get()

	util.AssertEqual(t, 1, r)
	util.AssertEqual(t, 2, c)
}

func TestPointAdd(t *testing.T) {
	point := c4.Point{1, 2}
	other := c4.Point{3, -4}

	result := point.Add(other)

	util.AssertEqual(t, c4.Point{4, -2}, result)
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

	util.AssertEqual(t, c4.Point{-1, 0}, n)
	util.AssertEqual(t, c4.Point{-1, 1}, ne)
	util.AssertEqual(t, c4.Point{0, 1}, e)
	util.AssertEqual(t, c4.Point{1, 1}, se)
	util.AssertEqual(t, c4.Point{1, 0}, s)
	util.AssertEqual(t, c4.Point{1, -1}, sw)
	util.AssertEqual(t, c4.Point{0, -1}, w)
	util.AssertEqual(t, c4.Point{-1, -1}, nw)
}
