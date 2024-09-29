package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func assertEqual(t *testing.T, c *util.Cursor, row, col int) {
	t.Helper()
	actual_row, actual_col := c.Get()
	util.AssertEqual(t, row, actual_row)
	util.AssertEqual(t, col, actual_col)
	util.AssertEqual(t, row, c.Row())
	util.AssertEqual(t, col, c.Col())
}

func TestNewCursor(t *testing.T) {
	c := util.NewCursor()
	assertEqual(t, c, 0, 0)
}

func TestCursorMoveUp(t *testing.T) {
	c := util.NewCursor()
	c.MoveUp()
	assertEqual(t, c, -1, 0)
}

func TestCursorMoveDown(t *testing.T) {
	c := util.NewCursor()
	c.MoveDown()
	assertEqual(t, c, 1, 0)
}

func TestCursorMoveLeft(t *testing.T) {
	c := util.NewCursor()
	c.MoveLeft()
	assertEqual(t, c, 0, -1)
}

func TestCursorMoveRight(t *testing.T) {
	c := util.NewCursor()
	c.MoveRight()
	assertEqual(t, c, 0, 1)
}

func TestCursorConstrainRowUp(t *testing.T) {
	c := util.NewCursor()
	c.MoveUp()

	// Minimum is inclusive
	c.ConstrainRow(-1, 0)
	assertEqual(t, c, -1, 0)
	c.ConstrainRow(0, 0)
	assertEqual(t, c, 0, 0) // min takes precedence over max (max is exclusive)
}

func TestCursorConstrainRowDown(t *testing.T) {
	c := util.NewCursor()
	c.MoveDown()

	// Maximum is exclusive
	c.ConstrainRow(0, 2)
	assertEqual(t, c, 1, 0)
	c.ConstrainRow(0, 1)
	assertEqual(t, c, 0, 0)
}

func TestCursorConstrainColLeft(t *testing.T) {
	c := util.NewCursor()
	c.MoveLeft()

	// Minimum is inclusive
	c.ConstrainCol(-1, 0)
	assertEqual(t, c, 0, -1)
	c.ConstrainCol(0, 0)
	assertEqual(t, c, 0, 0) // min takes precedence over max (max is exclusive)
}

func TestCursorConstrainColRight(t *testing.T) {
	c := util.NewCursor()
	c.MoveRight()

	// Maximum is exclusive
	c.ConstrainCol(0, 2)
	assertEqual(t, c, 0, 1)
	c.ConstrainCol(0, 1)
	assertEqual(t, c, 0, 0)
}
