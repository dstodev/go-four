package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestNewBoard(t *testing.T) {
	b := c4.NewBoard(1, 2)

	assertEqual(t, 1, b.Rows())
	assertEqual(t, 2, b.Columns())
	assertEqual(t, c4.None, b.Get(0, 0))
}

func TestBoardIndexing(t *testing.T) {
	b := c4.NewBoard(1, 1)

	// Get in- and out-of bounds
	assertEqual(t, c4.None, b.Get(0, 0))
	assertEqual(t, c4.None, b.Get(0, 1))
	assertEqual(t, c4.None, b.Get(0, -1))
	assertEqual(t, c4.None, b.Get(1, 0))
	assertEqual(t, c4.None, b.Get(-1, 0))

	// Set in- and out-of bounds
	b.
		Set(0, 0, c4.One).
		Set(0, 1, c4.One).
		Set(0, -1, c4.One).
		Set(1, 0, c4.One).
		Set(-1, 0, c4.One)

	// Assert only in-bounds changes
	assertEqual(t, c4.One, b.Get(0, 0))
	assertEqual(t, c4.None, b.Get(0, 1))
	assertEqual(t, c4.None, b.Get(0, -1))
	assertEqual(t, c4.None, b.Get(1, 0))
	assertEqual(t, c4.None, b.Get(-1, 0))
}

func TestBoardIsEqual(t *testing.T) {
	b1 := c4.NewBoard(1, 1)
	b2 := c4.NewBoard(1, 1)
	b3 := c4.NewBoard(1, 2)
	b4 := c4.NewBoard(2, 1)

	assertEqual(t, true, b1.IsEqual(b2))
	assertEqual(t, false, b1.IsEqual(b3))
	assertEqual(t, false, b1.IsEqual(b4))

	b1.Set(0, 0, c4.One)

	assertEqual(t, false, b1.IsEqual(b2))
}

func TestBoardClone(t *testing.T) {
	b1 := c4.NewBoard(1, 1)
	b2 := b1
	b3 := b1.Clone()

	b1.Set(0, 0, c4.One)

	assertEqual(t, c4.One, b2.Get(0, 0))
	assertEqual(t, c4.None, b3.Get(0, 0))
}

func TestBoardNeighbor(t *testing.T) {
	b := c4.NewBoard(3, 3)

	// Set up board
	// 0 1 2
	// 3 4 5
	// 6 7 8
	b.
		Set(0, 0, 0).
		Set(0, 1, 1).
		Set(0, 2, 2).
		Set(1, 0, 3).
		Set(1, 1, 4).
		Set(1, 2, 5).
		Set(2, 0, 6).
		Set(2, 1, 7).
		Set(2, 2, 8)

	// Test all directions centered on 4
	assertEqual(t, 1, b.Neighbor(1, 1, c4.North))
	assertEqual(t, 2, b.Neighbor(1, 1, c4.NorthEast))
	assertEqual(t, 5, b.Neighbor(1, 1, c4.East))
	assertEqual(t, 8, b.Neighbor(1, 1, c4.SouthEast))
	assertEqual(t, 7, b.Neighbor(1, 1, c4.South))
	assertEqual(t, 6, b.Neighbor(1, 1, c4.SouthWest))
	assertEqual(t, 3, b.Neighbor(1, 1, c4.West))
	assertEqual(t, 0, b.Neighbor(1, 1, c4.NorthWest))

	// Test out of bounds
	assertEqual(t, c4.None, b.Neighbor(0, 0, c4.North))
}

func TestBoardCountDirection(t *testing.T) {
	b := c4.NewBoard(5, 5)

	b.
		Set(0, 2, c4.One).
		Set(0, 4, c4.One).
		Set(1, 2, c4.One).
		Set(1, 3, c4.One).
		Set(2, 1, c4.One).
		Set(2, 2, c4.One).
		Set(2, 3, c4.Two).
		Set(2, 4, c4.One).
		Set(3, 1, c4.One).
		Set(3, 2, c4.One).
		Set(3, 3, c4.Two)

	/*
		   0 1 2 3 4
		0  - - R - R
		1  - - R R -
		2  - R R B R  <-- Test focuses on the center R at (2,2)
		3  - R R B -
		4  - - - - -  <-- and at the bottom-left (4,0)
	*/

	assertEqual(t, 3, b.CountDirection(2, 2, c4.North))
	assertEqual(t, 3, b.CountDirection(2, 2, c4.NorthEast))
	assertEqual(t, 1, b.CountDirection(2, 2, c4.East))
	assertEqual(t, 1, b.CountDirection(2, 2, c4.SouthEast))
	assertEqual(t, 2, b.CountDirection(2, 2, c4.South))
	assertEqual(t, 2, b.CountDirection(2, 2, c4.SouthWest))
	assertEqual(t, 2, b.CountDirection(2, 2, c4.West))
	assertEqual(t, 1, b.CountDirection(2, 2, c4.NorthWest))

	// At the empty bottom-left, facing top-right
	assertEqual(t, 1, b.CountDirection(4, 0, c4.NorthEast))

	// Out of bounds
	assertEqual(t, 0, b.CountDirection(-1, 0, c4.North))
}

func TestBoardCountBidirection(t *testing.T) {
	b := c4.NewBoard(5, 5)

	b.
		Set(0, 2, c4.One).
		Set(0, 4, c4.One).
		Set(1, 2, c4.One).
		Set(1, 3, c4.One).
		Set(2, 1, c4.One).
		Set(2, 2, c4.One).
		Set(2, 3, c4.Two).
		Set(2, 4, c4.One).
		Set(3, 1, c4.One).
		Set(3, 2, c4.One).
		Set(3, 3, c4.Two)

	/*
		   0 1 2 3 4
		0  - - R - R
		1  - - R R -
		2  - R R B R  <-- Test focuses on the center R at (2,2)
		3  - R R B -
		4  - - - - -  <-- and at the bottom-left (4,0)
	*/

	assertEqual(t, 4, b.CountBidirection(2, 2, c4.North))
	assertEqual(t, 4, b.CountBidirection(2, 2, c4.NorthEast))
	assertEqual(t, 2, b.CountBidirection(2, 2, c4.East))
	assertEqual(t, 1, b.CountBidirection(2, 2, c4.SouthEast))
	assertEqual(t, 4, b.CountBidirection(2, 2, c4.South))
	assertEqual(t, 4, b.CountBidirection(2, 2, c4.SouthWest))
	assertEqual(t, 2, b.CountBidirection(2, 2, c4.West))
	assertEqual(t, 1, b.CountBidirection(2, 2, c4.NorthWest))

	// At the empty bottom-left, facing top-right
	assertEqual(t, 1, b.CountBidirection(4, 0, c4.NorthEast))

	// Out of bounds
	assertEqual(t, 0, b.CountBidirection(-1, 0, c4.North))
}
