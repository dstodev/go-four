package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestDirectionString(t *testing.T) {
	assertEqual(t, "North", c4.North.String())
	assertEqual(t, "NorthEast", c4.NorthEast.String())
	assertEqual(t, "East", c4.East.String())
	assertEqual(t, "SouthEast", c4.SouthEast.String())
	assertEqual(t, "South", c4.South.String())
	assertEqual(t, "SouthWest", c4.SouthWest.String())
	assertEqual(t, "West", c4.West.String())
	assertEqual(t, "NorthWest", c4.NorthWest.String())
}

func TestDirectionNegate(t *testing.T) {
	assertEqual(t, c4.South, c4.North.Negate())
	assertEqual(t, c4.SouthWest, c4.NorthEast.Negate())
	assertEqual(t, c4.West, c4.East.Negate())
	assertEqual(t, c4.NorthWest, c4.SouthEast.Negate())
	assertEqual(t, c4.North, c4.South.Negate())
	assertEqual(t, c4.NorthEast, c4.SouthWest.Negate())
	assertEqual(t, c4.East, c4.West.Negate())
	assertEqual(t, c4.SouthEast, c4.NorthWest.Negate())
}

func TestDirectionOffsetRow(t *testing.T) {
	assertEqual(t, -1, c4.North.OffsetRow())
	assertEqual(t, -1, c4.NorthEast.OffsetRow())
	assertEqual(t, 0, c4.East.OffsetRow())
	assertEqual(t, 1, c4.SouthEast.OffsetRow())
	assertEqual(t, 1, c4.South.OffsetRow())
	assertEqual(t, 1, c4.SouthWest.OffsetRow())
	assertEqual(t, 0, c4.West.OffsetRow())
	assertEqual(t, -1, c4.NorthWest.OffsetRow())
}

func TestDirectionOffsetColumn(t *testing.T) {
	assertEqual(t, 0, c4.North.OffsetColumn())
	assertEqual(t, 1, c4.NorthEast.OffsetColumn())
	assertEqual(t, 1, c4.East.OffsetColumn())
	assertEqual(t, 1, c4.SouthEast.OffsetColumn())
	assertEqual(t, 0, c4.South.OffsetColumn())
	assertEqual(t, -1, c4.SouthWest.OffsetColumn())
	assertEqual(t, -1, c4.West.OffsetColumn())
	assertEqual(t, -1, c4.NorthWest.OffsetColumn())
}

func TestDirectionOffset(t *testing.T) {
	assertEqual(t, c4.Point{-1, 0}, c4.North.Offset())
	assertEqual(t, c4.Point{-1, 1}, c4.NorthEast.Offset())
	assertEqual(t, c4.Point{0, 1}, c4.East.Offset())
	assertEqual(t, c4.Point{1, 1}, c4.SouthEast.Offset())
	assertEqual(t, c4.Point{1, 0}, c4.South.Offset())
	assertEqual(t, c4.Point{1, -1}, c4.SouthWest.Offset())
	assertEqual(t, c4.Point{0, -1}, c4.West.Offset())
	assertEqual(t, c4.Point{-1, -1}, c4.NorthWest.Offset())
}
