package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestDirectionString(t *testing.T) {
	util.AssertEqual(t, "None", c4.NoDirection.String())
	util.AssertEqual(t, "North", c4.North.String())
	util.AssertEqual(t, "NorthEast", c4.NorthEast.String())
	util.AssertEqual(t, "East", c4.East.String())
	util.AssertEqual(t, "SouthEast", c4.SouthEast.String())
	util.AssertEqual(t, "South", c4.South.String())
	util.AssertEqual(t, "SouthWest", c4.SouthWest.String())
	util.AssertEqual(t, "West", c4.West.String())
	util.AssertEqual(t, "NorthWest", c4.NorthWest.String())
}

func TestDirectionDefault(t *testing.T) {
	var d c4.Direction
	util.AssertEqual(t, c4.NoDirection, d)
}

func TestDirectionNegate(t *testing.T) {
	util.AssertEqual(t, c4.NoDirection, c4.NoDirection.Negate())
	util.AssertEqual(t, c4.South, c4.North.Negate())
	util.AssertEqual(t, c4.SouthWest, c4.NorthEast.Negate())
	util.AssertEqual(t, c4.West, c4.East.Negate())
	util.AssertEqual(t, c4.NorthWest, c4.SouthEast.Negate())
	util.AssertEqual(t, c4.North, c4.South.Negate())
	util.AssertEqual(t, c4.NorthEast, c4.SouthWest.Negate())
	util.AssertEqual(t, c4.East, c4.West.Negate())
	util.AssertEqual(t, c4.SouthEast, c4.NorthWest.Negate())
}

func TestDirectionOffsetRow(t *testing.T) {
	util.AssertEqual(t, 0, c4.NoDirection.OffsetRow())
	util.AssertEqual(t, -1, c4.North.OffsetRow())
	util.AssertEqual(t, -1, c4.NorthEast.OffsetRow())
	util.AssertEqual(t, 0, c4.East.OffsetRow())
	util.AssertEqual(t, 1, c4.SouthEast.OffsetRow())
	util.AssertEqual(t, 1, c4.South.OffsetRow())
	util.AssertEqual(t, 1, c4.SouthWest.OffsetRow())
	util.AssertEqual(t, 0, c4.West.OffsetRow())
	util.AssertEqual(t, -1, c4.NorthWest.OffsetRow())
}

func TestDirectionOffsetColumn(t *testing.T) {
	util.AssertEqual(t, 0, c4.NoDirection.OffsetColumn())
	util.AssertEqual(t, 0, c4.North.OffsetColumn())
	util.AssertEqual(t, 1, c4.NorthEast.OffsetColumn())
	util.AssertEqual(t, 1, c4.East.OffsetColumn())
	util.AssertEqual(t, 1, c4.SouthEast.OffsetColumn())
	util.AssertEqual(t, 0, c4.South.OffsetColumn())
	util.AssertEqual(t, -1, c4.SouthWest.OffsetColumn())
	util.AssertEqual(t, -1, c4.West.OffsetColumn())
	util.AssertEqual(t, -1, c4.NorthWest.OffsetColumn())
}

func TestDirectionOffset(t *testing.T) {
	util.AssertEqual(t, c4.Point{0, 0}, c4.NoDirection.Offset())
	util.AssertEqual(t, c4.Point{-1, 0}, c4.North.Offset())
	util.AssertEqual(t, c4.Point{-1, 1}, c4.NorthEast.Offset())
	util.AssertEqual(t, c4.Point{0, 1}, c4.East.Offset())
	util.AssertEqual(t, c4.Point{1, 1}, c4.SouthEast.Offset())
	util.AssertEqual(t, c4.Point{1, 0}, c4.South.Offset())
	util.AssertEqual(t, c4.Point{1, -1}, c4.SouthWest.Offset())
	util.AssertEqual(t, c4.Point{0, -1}, c4.West.Offset())
	util.AssertEqual(t, c4.Point{-1, -1}, c4.NorthWest.Offset())
}
