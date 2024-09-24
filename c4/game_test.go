package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestNewGame(t *testing.T) {
	game := c4.NewGame()

	assertEqual(t, c4.Initial, game.Status())
	assertEqual(t, c4.None, game.Turn())
	assertEqual(t, 0, game.TurnCount())

	assertEqual(t, 6, game.Board().RowCount())
	assertEqual(t, 7, game.Board().ColCount())
	assertEqual(t, 4, game.ToWin())
	assertEqual(t, 0, game.MaxTurns())
}

func TestNewGameSize(t *testing.T) {
	game := c4.NewGame(2, 3)

	assertEqual(t, 2, game.Board().RowCount())
	assertEqual(t, 3, game.Board().ColCount())
	assertEqual(t, 4, game.ToWin())
	assertEqual(t, 0, game.MaxTurns())
}

func TestNewGameToWin(t *testing.T) {
	expected := c4.NewGame().ToWin() + 1
	game := c4.NewGame(2, 3, expected)

	assertEqual(t, 2, game.Board().RowCount())
	assertEqual(t, 3, game.Board().ColCount())
	assertEqual(t, expected, game.ToWin())
	assertEqual(t, 0, game.MaxTurns())
}

func TestNewGameMaxTurns(t *testing.T) {
	expected := c4.NewGame().MaxTurns() + 1
	game := c4.NewGame(2, 3, 4, expected)

	assertEqual(t, 2, game.Board().RowCount())
	assertEqual(t, 3, game.Board().ColCount())
	assertEqual(t, 4, game.ToWin())
	assertEqual(t, expected, game.MaxTurns())
}

func TestGameStart(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 0, game.TurnCount())
}

func TestGamePlayTurnBeforeStart(t *testing.T) {
	game := c4.NewGame()
	board := game.Board().Clone()
	game.PlayTurn(0)

	assertEqual(t, true, board.IsEqual(game.Board()))
	assertEqual(t, c4.Initial, game.Status())
	assertEqual(t, c4.None, game.Turn())
	assertEqual(t, 0, game.TurnCount())
}

func TestGamePlayTurnOutOfBounds(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	board := game.Board().Clone()

	game.PlayTurn(-1)

	assertEqual(t, c4.One, game.Turn())

	game.PlayTurn(game.Board().ColCount())

	assertEqual(t, true, board.IsEqual(game.Board()))
	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 0, game.TurnCount())
}

func TestGameTurnCount(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 0, game.TurnCount())

	game.PlayTurn(0)

	assertEqual(t, c4.Two, game.Turn())
	assertEqual(t, 1, game.TurnCount())

	game.PlayTurn(1)

	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 2, game.TurnCount())
}

func TestGamePlayTurnCol0(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)

	/*
	     0 1 2 3 4 5 6
	   0 - - - - - - -
	   1 - - - - - - -
	   2 - - - - - - -
	   3 - - - - - - -
	   4 - - - - - - -
	   5 A - - - - - -
	*/

	assertEqual(t, c4.One, game.Board().Get(5, 0))
	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.Two, game.Turn())
}

func TestGameFillCol0(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)

	/*
	     0 1 2 3 4 5 6
	   0 B - - - - - -
	   1 A - - - - - -
	   2 B - - - - - -
	   3 A - - - - - -
	   4 B - - - - - -
	   5 A - - - - - -
	*/

	assertEqual(t, c4.One, game.Board().Get(5, 0))
	assertEqual(t, c4.Two, game.Board().Get(4, 0))
	assertEqual(t, c4.One, game.Board().Get(3, 0))
	assertEqual(t, c4.Two, game.Board().Get(2, 0))
	assertEqual(t, c4.One, game.Board().Get(1, 0))
	assertEqual(t, c4.Two, game.Board().Get(0, 0))

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())

	// Try to overfill the column
	board := game.Board().Clone()
	game.PlayTurn(0)

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, true, board.IsEqual(game.Board()))
}

func TestFullGamePlayerOneWinsVertical(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)

	/*
	     0 1 2 3 4 5 6     0 1 2 3 4 5 6
	   0 - - - - - - -   0 - - - - - - -
	   1 - - - - - - -   1 - - - - - - -
	   2 - - - - - - -   2 R - - - - - -
	   3 R B - - - - -   3 R B - - - - -
	   4 R B - - - - -   4 R B - - - - -
	   5 R B - - - - -   5 R B - - - - -
	*/

	assertEqual(t, c4.One, game.Board().Get(5, 0))
	assertEqual(t, c4.One, game.Board().Get(4, 0))
	assertEqual(t, c4.One, game.Board().Get(3, 0))
	assertEqual(t, c4.Two, game.Board().Get(5, 1))
	assertEqual(t, c4.Two, game.Board().Get(4, 1))
	assertEqual(t, c4.Two, game.Board().Get(3, 1))

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 6, game.TurnCount())

	game.PlayTurn(0)

	assertEqual(t, c4.One, game.Board().Get(2, 0))

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 7, game.TurnCount())
}

func TestCannotPlayTurnAfterVictory(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())

	board := game.Board().Clone()
	game.PlayTurn(2)

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, true, board.IsEqual(game.Board()))
}

func TestFullGamePlayerOneWinsLongHorizontal(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(1)
	game.PlayTurn(2)
	game.PlayTurn(2)
	game.PlayTurn(4)
	game.PlayTurn(4)

	/*
	     0 1 2 3 4 5 6     0 1 2 3 4 5 6
	   0 - - - - - - -   0 - - - - - - -
	   1 - - - - - - -   1 - - - - - - -
	   2 - - - - - - -   2 - - - - - - -
	   3 - - - - - - -   3 - - - - - - -
	   4 B B B - B - -   4 B B B - B - -
	   5 A A A - A - -   5 A A A A A - -
	*/

	assertEqual(t, c4.One, game.Board().Get(5, 0))
	assertEqual(t, c4.One, game.Board().Get(5, 1))
	assertEqual(t, c4.One, game.Board().Get(5, 2))
	assertEqual(t, c4.One, game.Board().Get(5, 4))
	assertEqual(t, c4.Two, game.Board().Get(4, 0))
	assertEqual(t, c4.Two, game.Board().Get(4, 1))
	assertEqual(t, c4.Two, game.Board().Get(4, 2))
	assertEqual(t, c4.Two, game.Board().Get(4, 4))

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 8, game.TurnCount())

	game.PlayTurn(3)

	assertEqual(t, c4.One, game.Board().Get(5, 3))

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 9, game.TurnCount())
}

func TestFullGamePlayerOneWinsDiagonal(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(1)
	game.PlayTurn(2)
	game.PlayTurn(3)
	game.PlayTurn(2)
	game.PlayTurn(3)
	game.PlayTurn(3)
	game.PlayTurn(3)
	game.PlayTurn(4)

	/*
	     0 1 2 3 4 5 6     0 1 2 3 4 5 6
	   0 - - - - - - -   0 - - - - - - -
	   1 - - - - - - -   1 - - - - - - -
	   2 - - - A - - -   2 - - - A - - -
	   3 - - - B - - -   3 - - A B - - -
	   4 - A B A - - -   4 - A B A - - -
	   5 A B B A B - -   5 A B B A B - -
	*/

	assertEqual(t, c4.One, game.Board().Get(5, 0))
	assertEqual(t, c4.Two, game.Board().Get(5, 1))
	assertEqual(t, c4.Two, game.Board().Get(5, 2))
	assertEqual(t, c4.One, game.Board().Get(5, 3))
	assertEqual(t, c4.Two, game.Board().Get(5, 4))
	assertEqual(t, c4.One, game.Board().Get(4, 1))
	assertEqual(t, c4.Two, game.Board().Get(4, 2))
	assertEqual(t, c4.One, game.Board().Get(4, 3))
	assertEqual(t, c4.Two, game.Board().Get(3, 3))
	assertEqual(t, c4.One, game.Board().Get(2, 3))

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 10, game.TurnCount())

	game.PlayTurn(2)

	assertEqual(t, c4.One, game.Board().Get(3, 2))

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 11, game.TurnCount())
}

func TestFullGameNoWinner(t *testing.T) {
	game := c4.NewGame(1, 2)
	game.Start()
	game.PlayTurn(0)

	/*
	     0 1     0 1
	   0 A -   0 A B
	*/

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.Two, game.Turn())
	assertEqual(t, 1, game.TurnCount())

	game.PlayTurn(1)

	assertEqual(t, c4.Draw, game.Status())
	assertEqual(t, c4.None, game.Turn())
	assertEqual(t, 2, game.TurnCount())
}

func TestFullGameLastWinner(t *testing.T) {
	game := c4.NewGame(4, 4)
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(2)
	game.PlayTurn(3)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(2)
	game.PlayTurn(3)
	game.PlayTurn(0)
	game.PlayTurn(2)
	game.PlayTurn(1)
	game.PlayTurn(3)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(2)

	/*
	     0 1 2 3     0 1 2 3
	   0 A B A -   0 A B A B
	   1 A A B B   1 A A B B
	   2 B A A B   2 B A A B
	   3 A B A B   3 A B A B
	*/

	assertEqual(t, c4.One, game.Board().Get(3, 0))
	assertEqual(t, c4.Two, game.Board().Get(3, 1))
	assertEqual(t, c4.One, game.Board().Get(3, 2))
	assertEqual(t, c4.Two, game.Board().Get(3, 3))
	assertEqual(t, c4.Two, game.Board().Get(2, 0))
	assertEqual(t, c4.One, game.Board().Get(2, 1))
	assertEqual(t, c4.One, game.Board().Get(2, 2))
	assertEqual(t, c4.Two, game.Board().Get(2, 3))
	assertEqual(t, c4.One, game.Board().Get(1, 0))
	assertEqual(t, c4.One, game.Board().Get(1, 1))
	assertEqual(t, c4.Two, game.Board().Get(1, 2))
	assertEqual(t, c4.Two, game.Board().Get(1, 3))
	assertEqual(t, c4.One, game.Board().Get(0, 0))
	assertEqual(t, c4.Two, game.Board().Get(0, 1))
	assertEqual(t, c4.One, game.Board().Get(0, 2))

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.Two, game.Turn())
	assertEqual(t, 15, game.TurnCount())

	game.PlayTurn(3)

	assertEqual(t, c4.Two, game.Board().Get(0, 3))

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.Two, game.Turn())
	assertEqual(t, 16, game.TurnCount())
}

func TestGameStartResetsGame(t *testing.T) {
	game := c4.NewGame()

	game.Start()
	game.PlayTurn(0)

	game.Start()

	empty := c4.NewGame()
	empty.Start()

	assertEqual(t, empty.Status(), game.Status())
	assertEqual(t, empty.Turn(), game.Turn())
	assertEqual(t, empty.TurnCount(), game.TurnCount())
	assertEqual(t, true, empty.Board().IsEqual(game.Board()))
}

func TestFullGamePlayerOneWins5InARow(t *testing.T) {
	game := c4.NewGame(6, 7, 5)
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(1)
	game.PlayTurn(2)
	game.PlayTurn(2)
	game.PlayTurn(3)
	game.PlayTurn(3)

	/*
	     0 1 2 3 4 5 6     0 1 2 3 4 5 6
	   0 - - - - - - -   0 - - - - - - -
	   1 - - - - - - -   1 - - - - - - -
	   2 - - - - - - -   2 - - - - - - -
	   3 - - - - - - -   3 - - - - - - -
	   4 B B B B - - -   4 B B B B - - -
	   5 A A A A - - -   5 A A A A A - -
	*/

	assertEqual(t, c4.One, game.Board().Get(5, 0))
	assertEqual(t, c4.Two, game.Board().Get(4, 0))
	assertEqual(t, c4.One, game.Board().Get(5, 1))
	assertEqual(t, c4.Two, game.Board().Get(4, 1))
	assertEqual(t, c4.One, game.Board().Get(5, 2))
	assertEqual(t, c4.Two, game.Board().Get(4, 2))
	assertEqual(t, c4.One, game.Board().Get(5, 3))
	assertEqual(t, c4.Two, game.Board().Get(4, 3))

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 8, game.TurnCount())

	game.PlayTurn(4)

	assertEqual(t, c4.One, game.Board().Get(5, 4))

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 9, game.TurnCount())
}

func TestGameMaxTurnDraw1(t *testing.T) {
	game := c4.NewGame(6, 7, 4, 1)
	game.Start()

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 0, game.TurnCount())

	game.PlayTurn(0)

	assertEqual(t, c4.Draw, game.Status())
	assertEqual(t, c4.None, game.Turn())
	assertEqual(t, 1, game.TurnCount())
}

func TestGameMaxTurnDraw3(t *testing.T) {
	game := c4.NewGame(6, 7, 2, 3)
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 2, game.TurnCount())

	game.PlayTurn(2)

	assertEqual(t, c4.Draw, game.Status())
	assertEqual(t, c4.None, game.Turn())
	assertEqual(t, 3, game.TurnCount())
}

func TestGameWinMaxTurn3(t *testing.T) {
	game := c4.NewGame(6, 7, 2, 3)
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)

	assertEqual(t, c4.Running, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 2, game.TurnCount())

	game.PlayTurn(1)

	assertEqual(t, c4.Completed, game.Status())
	assertEqual(t, c4.One, game.Turn())
	assertEqual(t, 3, game.TurnCount())
}
