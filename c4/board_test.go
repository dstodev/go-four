package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestNewBoard(t *testing.T) {
	b := c4.NewBoard(1, 2)

	util.AssertEqual(t, 1, b.RowCount())
	util.AssertEqual(t, 2, b.ColCount())
	util.AssertEqual(t, c4.None, b.Get(0, 0))
}

func TestBoardInBounds(t *testing.T) {
	b := c4.NewBoard(1, 1)

	util.AssertEqual(t, true, b.InBounds(0, 0))
	util.AssertEqual(t, false, b.InBounds(-1, 1))
	util.AssertEqual(t, false, b.InBounds(-1, 0))
	util.AssertEqual(t, false, b.InBounds(-1, 1))
	util.AssertEqual(t, false, b.InBounds(0, -1))
	util.AssertEqual(t, false, b.InBounds(0, 1))
	util.AssertEqual(t, false, b.InBounds(1, -1))
	util.AssertEqual(t, false, b.InBounds(1, 0))
	util.AssertEqual(t, false, b.InBounds(1, 1))
}

func TestBoardIndexing(t *testing.T) {
	b := c4.NewBoard(1, 1)

	// Get in- and out-of bounds
	util.AssertEqual(t, c4.None, b.Get(0, 0))
	util.AssertEqual(t, c4.None, b.Get(0, 1))
	util.AssertEqual(t, c4.None, b.Get(0, -1))
	util.AssertEqual(t, c4.None, b.Get(1, 0))
	util.AssertEqual(t, c4.None, b.Get(-1, 0))

	// Set in- and out-of bounds
	b.
		Set(0, 0, c4.One).
		Set(0, 1, c4.One).
		Set(0, -1, c4.One).
		Set(1, 0, c4.One).
		Set(-1, 0, c4.One)

	// Assert only in-bounds changes
	util.AssertEqual(t, c4.One, b.Get(0, 0))
	util.AssertEqual(t, c4.None, b.Get(0, 1))
	util.AssertEqual(t, c4.None, b.Get(0, -1))
	util.AssertEqual(t, c4.None, b.Get(1, 0))
	util.AssertEqual(t, c4.None, b.Get(-1, 0))
}

func TestBoardIsEqual(t *testing.T) {
	b1 := c4.NewBoard(1, 1)
	b2 := c4.NewBoard(1, 1)
	b3 := c4.NewBoard(1, 2)
	b4 := c4.NewBoard(2, 1)

	util.AssertEqual(t, true, b1.IsEqual(b2))
	util.AssertEqual(t, false, b1.IsEqual(b3))
	util.AssertEqual(t, false, b1.IsEqual(b4))

	b1.Set(0, 0, c4.One)

	util.AssertEqual(t, false, b1.IsEqual(b2))
}

func TestBoardClone(t *testing.T) {
	b1 := c4.NewBoard(1, 1)
	b2 := b1
	b3 := b1.Clone()

	b1.Set(0, 0, c4.One)

	util.AssertEqual(t, c4.One, b2.Get(0, 0))
	util.AssertEqual(t, c4.None, b3.Get(0, 0))
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
	util.AssertEqual(t, c4.Player(4), b.Neighbor(1, 1, c4.NoDirection))
	util.AssertEqual(t, c4.Player(1), b.Neighbor(1, 1, c4.North))
	util.AssertEqual(t, c4.Player(2), b.Neighbor(1, 1, c4.NorthEast))
	util.AssertEqual(t, c4.Player(5), b.Neighbor(1, 1, c4.East))
	util.AssertEqual(t, c4.Player(8), b.Neighbor(1, 1, c4.SouthEast))
	util.AssertEqual(t, c4.Player(7), b.Neighbor(1, 1, c4.South))
	util.AssertEqual(t, c4.Player(6), b.Neighbor(1, 1, c4.SouthWest))
	util.AssertEqual(t, c4.Player(3), b.Neighbor(1, 1, c4.West))
	util.AssertEqual(t, c4.Player(0), b.Neighbor(1, 1, c4.NorthWest))

	// Test out of bounds
	util.AssertEqual(t, c4.None, b.Neighbor(0, 0, c4.North))
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
	   0 - - A - A
	   1 - - A A -
	   2 - A A B A  <-- Test focuses on the center R at (2,2)
	   3 - A A B -
	   4 - - - - -  <-- and at the bottom-left (4,0)
	*/

	util.AssertEqual(t, 3, b.CountDirection(2, 2, c4.North))
	util.AssertEqual(t, 3, b.CountDirection(2, 2, c4.NorthEast))
	util.AssertEqual(t, 1, b.CountDirection(2, 2, c4.East))
	util.AssertEqual(t, 1, b.CountDirection(2, 2, c4.SouthEast))
	util.AssertEqual(t, 2, b.CountDirection(2, 2, c4.South))
	util.AssertEqual(t, 2, b.CountDirection(2, 2, c4.SouthWest))
	util.AssertEqual(t, 2, b.CountDirection(2, 2, c4.West))
	util.AssertEqual(t, 1, b.CountDirection(2, 2, c4.NorthWest))

	// At the empty bottom-left, facing top-right
	util.AssertEqual(t, 1, b.CountDirection(4, 0, c4.NorthEast))

	// Out of bounds
	util.AssertEqual(t, 0, b.CountDirection(-1, 0, c4.North))
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
	   0 - - A - A
	   1 - - A A -
	   2 - A A B A  <-- Test focuses on the center R at (2,2)
	   3 - A A B -
	   4 - - - - -  <-- and at the bottom-left (4,0)
	*/

	util.AssertEqual(t, 4, b.CountBidirection(2, 2, c4.North))
	util.AssertEqual(t, 4, b.CountBidirection(2, 2, c4.NorthEast))
	util.AssertEqual(t, 2, b.CountBidirection(2, 2, c4.East))
	util.AssertEqual(t, 1, b.CountBidirection(2, 2, c4.SouthEast))
	util.AssertEqual(t, 4, b.CountBidirection(2, 2, c4.South))
	util.AssertEqual(t, 4, b.CountBidirection(2, 2, c4.SouthWest))
	util.AssertEqual(t, 2, b.CountBidirection(2, 2, c4.West))
	util.AssertEqual(t, 1, b.CountBidirection(2, 2, c4.NorthWest))

	// At the empty bottom-left, facing top-right
	// Counts player 0 ("None") as well, hence 1 where there is no player
	util.AssertEqual(t, 1, b.CountBidirection(4, 0, c4.NorthEast))

	// Out of bounds
	util.AssertEqual(t, 0, b.CountBidirection(-1, 2, c4.North))
	util.AssertEqual(t, 0, b.CountBidirection(2, 5, c4.West))
}

func TestBoardRows(t *testing.T) {
	b := c4.NewBoard(3, 4)

	// Set up board
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 a b
	b.
		Set(0, 0, 0).
		Set(0, 1, 1).
		Set(0, 2, 2).
		Set(0, 3, 3).
		Set(1, 0, 4).
		Set(1, 1, 5).
		Set(1, 2, 6).
		Set(1, 3, 7).
		Set(2, 0, 8).
		Set(2, 1, 9).
		Set(2, 2, 10).
		Set(2, 3, 11)

	for i, row := range b.Rows() {
		for j, cell := range row {
			util.AssertEqual(t, i*b.ColCount()+j, int(cell))
		}
	}

}
