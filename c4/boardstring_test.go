package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestToBoardStringNone(t *testing.T) {
	game := c4.NewGame(1, 1)

	expected := ""
	actual := c4.ToBoardString(game)

	util.AssertEqual(t, expected, actual)
}
func TestToBoardStringOne(t *testing.T) {
	game := c4.NewGame(1, 1)
	game.Start()

	game.PlayTurn(0)

	expected := "0"
	actual := c4.ToBoardString(game)

	util.AssertEqual(t, expected, actual)
}

func TestToBoardStringTwo(t *testing.T) {
	game := c4.NewGame(1, 2)
	game.Start()

	game.PlayTurn(0)
	game.PlayTurn(1)

	expected := "0,1"
	actual := c4.ToBoardString(game)

	util.AssertEqual(t, expected, actual)
}

func TestFromBoardStringNone(t *testing.T) {
	game := c4.NewGame(6, 7)
	game = c4.FromBoardString(game, "")

	util.AssertEqual(t, c4.Initial, game.Status())
}

func TestFromBoardStringOne(t *testing.T) {
	game := c4.NewGame(6, 7)
	game = c4.FromBoardString(game, "0")

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, []int{0}, game.History())
}

func TestFromBoardStringTwo(t *testing.T) {
	game := c4.NewGame(6, 7)
	game = c4.FromBoardString(game, "0,1")

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(5, 1))
	util.AssertEqual(t, []int{0, 1}, game.History())
}

func TestFromBoardStringInvalid(t *testing.T) {
	game := c4.NewGame(6, 7)
	game = c4.FromBoardString(game, "a")

	util.AssertEqual(t, 0, game.TurnCount())

	game = c4.FromBoardString(game, "0,a,1")

	util.AssertEqual(t, 2, game.TurnCount())
}
