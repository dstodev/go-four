package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestNewGame(t *testing.T) {
	game := c4.NewGame()

	util.AssertEqual(t, c4.Initial, game.Status())
	util.AssertEqual(t, c4.None, game.Turn())
	util.AssertEqual(t, []int{}, game.History())
	util.AssertEqual(t, 0, game.TurnCount())

	util.AssertEqual(t, 6, game.Board().RowCount())
	util.AssertEqual(t, 7, game.Board().ColCount())
	util.AssertEqual(t, 4, game.ToWin())
	util.AssertEqual(t, 0, game.MaxTurns())
}

func TestNewGameSize(t *testing.T) {
	game := c4.NewGame(2, 3)

	util.AssertEqual(t, 2, game.Board().RowCount())
	util.AssertEqual(t, 3, game.Board().ColCount())
	util.AssertEqual(t, 4, game.ToWin())
	util.AssertEqual(t, 0, game.MaxTurns())
}

func TestNewGameToWin(t *testing.T) {
	expected := c4.NewGame().ToWin() + 1
	game := c4.NewGame(2, 3, expected)

	util.AssertEqual(t, 2, game.Board().RowCount())
	util.AssertEqual(t, 3, game.Board().ColCount())
	util.AssertEqual(t, expected, game.ToWin())
	util.AssertEqual(t, 0, game.MaxTurns())
}

func TestNewGameMaxTurns(t *testing.T) {
	expected := c4.NewGame().MaxTurns() + 1
	game := c4.NewGame(2, 3, 4, expected)

	util.AssertEqual(t, 2, game.Board().RowCount())
	util.AssertEqual(t, 3, game.Board().ColCount())
	util.AssertEqual(t, 4, game.ToWin())
	util.AssertEqual(t, expected, game.MaxTurns())
}

func TestGameStart(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 0, game.TurnCount())
}

func TestGameStartResetsGame(t *testing.T) {
	game := c4.NewGame()

	game.Start()
	game.PlayTurn(0)

	game.Start()

	empty := c4.NewGame()
	empty.Start()

	util.AssertEqual(t, empty.Status(), game.Status())
	util.AssertEqual(t, empty.Turn(), game.Turn())
	util.AssertEqual(t, empty.TurnCount(), game.TurnCount())
	util.AssertEqual(t, true, empty.Board().IsEqual(game.Board()))
	util.AssertEqual(t, empty.History(), game.History())
}

func TestGamePlayTurnBeforeStart(t *testing.T) {
	game := c4.NewGame()
	board := game.Board().Clone()
	game.PlayTurn(0)

	util.AssertEqual(t, true, board.IsEqual(game.Board()))
	util.AssertEqual(t, c4.Initial, game.Status())
	util.AssertEqual(t, c4.None, game.Turn())
	util.AssertEqual(t, 0, game.TurnCount())
}

func TestGamePlayTurnOutOfBounds(t *testing.T) {
	game := c4.NewGame()
	game.Start()
	board := game.Board().Clone()
	history := game.History()

	game.PlayTurn(-1)

	util.AssertEqual(t, c4.One, game.Turn())

	game.PlayTurn(game.Board().ColCount())

	util.AssertEqual(t, true, board.IsEqual(game.Board()))
	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, history, game.History())
	util.AssertEqual(t, 0, game.TurnCount())
}

func TestGameTurnCount(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 0, game.TurnCount())

	game.PlayTurn(0)

	util.AssertEqual(t, c4.Two, game.Turn())
	util.AssertEqual(t, 1, game.TurnCount())

	game.PlayTurn(1)

	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 2, game.TurnCount())
}

func TestGameHistory(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	util.AssertEqual(t, []int{}, game.History())

	game.PlayTurn(0)
	util.AssertEqual(t, []int{0}, game.History())

	game.PlayTurn(1)
	util.AssertEqual(t, []int{0, 1}, game.History())
}

func TestGameHistoryInitial(t *testing.T) {
	game := c4.NewGame()

	util.AssertEqual(t, []int{}, game.History())

	game.PlayTurn(0)
	util.AssertEqual(t, []int{}, game.History())
}

func TestGameHistoryCompleted(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)

	util.AssertEqual(t, []int{0, 1, 0, 1, 0, 1}, game.History())

	util.AssertEqual(t, c4.Running, game.Status())
	game.PlayTurn(0)
	util.AssertEqual(t, c4.Completed, game.Status())

	util.AssertEqual(t, []int{0, 1, 0, 1, 0, 1, 0}, game.History())

	game.PlayTurn(0)

	util.AssertEqual(t, []int{0, 1, 0, 1, 0, 1, 0}, game.History())
}

func TestGameHistoryInvalidPlacement(t *testing.T) {
	game := c4.NewGame()
	game.Start()

	// out of bounds
	game.PlayTurn(-1)
	game.PlayTurn(game.Board().ColCount())

	// column full
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(0)

	util.AssertEqual(t, 6, game.TurnCount())
	util.AssertEqual(t, []int{0, 0, 0, 0, 0, 0}, game.History())

	game.PlayTurn(0)

	util.AssertEqual(t, 6, game.TurnCount())
	util.AssertEqual(t, []int{0, 0, 0, 0, 0, 0}, game.History())
}

func TestGameHistoryDraw(t *testing.T) {
	game := c4.NewGame(1, 1)
	game.Start()

	game.PlayTurn(0)
	util.AssertEqual(t, []int{0}, game.History())

	game.PlayTurn(0)
	util.AssertEqual(t, []int{0}, game.History())
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

	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.Two, game.Turn())
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

	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(3, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(2, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(1, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(0, 0))

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())

	// Try to overfill the column
	board := game.Board().Clone()
	game.PlayTurn(0)

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, true, board.IsEqual(game.Board()))
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

	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(4, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(3, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(5, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(3, 1))

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 6, game.TurnCount())

	game.PlayTurn(0)

	util.AssertEqual(t, c4.One, game.Board().Get(2, 0))

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 7, game.TurnCount())
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

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())

	board := game.Board().Clone()
	game.PlayTurn(2)

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, true, board.IsEqual(game.Board()))
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

	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 1))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 2))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 4))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 2))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 4))

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 8, game.TurnCount())

	game.PlayTurn(3)

	util.AssertEqual(t, c4.One, game.Board().Get(5, 3))

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 9, game.TurnCount())
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

	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(5, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(5, 2))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 3))
	util.AssertEqual(t, c4.Two, game.Board().Get(5, 4))
	util.AssertEqual(t, c4.One, game.Board().Get(4, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 2))
	util.AssertEqual(t, c4.One, game.Board().Get(4, 3))
	util.AssertEqual(t, c4.Two, game.Board().Get(3, 3))
	util.AssertEqual(t, c4.One, game.Board().Get(2, 3))

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 10, game.TurnCount())

	game.PlayTurn(2)

	util.AssertEqual(t, c4.One, game.Board().Get(3, 2))

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 11, game.TurnCount())
}

func TestFullGameNoWinner(t *testing.T) {
	game := c4.NewGame(1, 2)
	game.Start()
	game.PlayTurn(0)

	/*
	     0 1     0 1
	   0 A -   0 A B
	*/

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.Two, game.Turn())
	util.AssertEqual(t, 1, game.TurnCount())

	game.PlayTurn(1)

	util.AssertEqual(t, c4.Draw, game.Status())
	util.AssertEqual(t, c4.None, game.Turn())
	util.AssertEqual(t, 2, game.TurnCount())
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

	util.AssertEqual(t, c4.One, game.Board().Get(3, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(3, 1))
	util.AssertEqual(t, c4.One, game.Board().Get(3, 2))
	util.AssertEqual(t, c4.Two, game.Board().Get(3, 3))
	util.AssertEqual(t, c4.Two, game.Board().Get(2, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(2, 1))
	util.AssertEqual(t, c4.One, game.Board().Get(2, 2))
	util.AssertEqual(t, c4.Two, game.Board().Get(2, 3))
	util.AssertEqual(t, c4.One, game.Board().Get(1, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(1, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(1, 2))
	util.AssertEqual(t, c4.Two, game.Board().Get(1, 3))
	util.AssertEqual(t, c4.One, game.Board().Get(0, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(0, 1))
	util.AssertEqual(t, c4.One, game.Board().Get(0, 2))

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.Two, game.Turn())
	util.AssertEqual(t, 15, game.TurnCount())

	game.PlayTurn(3)

	util.AssertEqual(t, c4.Two, game.Board().Get(0, 3))

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.Two, game.Turn())
	util.AssertEqual(t, 16, game.TurnCount())
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

	util.AssertEqual(t, c4.One, game.Board().Get(5, 0))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 0))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 1))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 1))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 2))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 2))
	util.AssertEqual(t, c4.One, game.Board().Get(5, 3))
	util.AssertEqual(t, c4.Two, game.Board().Get(4, 3))

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 8, game.TurnCount())

	game.PlayTurn(4)

	util.AssertEqual(t, c4.One, game.Board().Get(5, 4))

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 9, game.TurnCount())
}

func TestGameMaxTurnDraw1(t *testing.T) {
	game := c4.NewGame(6, 7, 4, 1)
	game.Start()

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 0, game.TurnCount())

	game.PlayTurn(0)

	util.AssertEqual(t, c4.Draw, game.Status())
	util.AssertEqual(t, c4.None, game.Turn())
	util.AssertEqual(t, 1, game.TurnCount())
}

func TestGameMaxTurnDraw3(t *testing.T) {
	game := c4.NewGame(6, 7, 2, 3)
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 2, game.TurnCount())

	game.PlayTurn(2)

	util.AssertEqual(t, c4.Draw, game.Status())
	util.AssertEqual(t, c4.None, game.Turn())
	util.AssertEqual(t, 3, game.TurnCount())
}

func TestGameWinMaxTurn3(t *testing.T) {
	game := c4.NewGame(6, 7, 2, 3)
	game.Start()
	game.PlayTurn(0)
	game.PlayTurn(1)

	util.AssertEqual(t, c4.Running, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 2, game.TurnCount())

	game.PlayTurn(1)

	util.AssertEqual(t, c4.Completed, game.Status())
	util.AssertEqual(t, c4.One, game.Turn())
	util.AssertEqual(t, 3, game.TurnCount())
}
