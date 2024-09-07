package c4

type Game struct {
	status GameStatus
	turn   Player
	board  Board
}

func NewGame() *Game {
	return &Game{
		status: Initial,
		turn:   None,
		board:  NewBoard(6, 7),
	}
}

func (game Game) Status() GameStatus {
	return game.status
}

func (game Game) Turn() Player {
	return game.turn
}

func (game Game) Board() Board {
	return game.board
}

func (game *Game) Start() {
	game.status = Running
	game.turn = One
}

func (game *Game) PlayTurn(column int) {
	if game.status != Running {
		return
	}

	if !game.board.inBounds(0, column) {
		return
	}

	for row := game.board.Rows() - 1; row >= 0; row-- {
		if game.board.Get(row, column) == None {
			game.board.Set(row, column, game.turn)
			break
		}
	}

	game.turn = game.turn.Negate()
}
