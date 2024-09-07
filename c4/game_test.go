package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestNewGame(t *testing.T) {
	game := c4.NewGame()

	assertEqual(t, c4.Initial, game.Status())
	assertEqual(t, c4.None, game.Turn())
}

func TestGameStart(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
}

func TestGamePlayTurnBeforeStart(t *testing.T) {
	game := c4.NewGame()
	board := game.Board().Clone()
	game.PlayTurn(0)

	assertEqual(t, c4.Initial, game.Status())
	assertEqual(t, c4.None, game.Turn())
	assertEqual(t, true, board.IsEqual(game.Board()))
}

func TestGamePlayTurnOutOfBounds(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	board := game.Board().Clone()

	game.PlayTurn(-1)

	assertEqual(t, c4.One, game.Turn())

	game.PlayTurn(game.Board().Columns())

	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, true, board.IsEqual(game.Board()))
}

func TestGamePlayTurnCol0(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)

	/*
	      0 1 2 3 4 5 6
	   0  - - - - - - -
	   1  - - - - - - -
	   2  - - - - - - -
	   3  - - - - - - -
	   4  - - - - - - -
	   5  R - - - - - -
	*/

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.Two, game.Turn())
	assertEqual(t, c4.One, game.Board().Get(5, 0))
}
