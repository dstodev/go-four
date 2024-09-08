package c4

type Game struct {
	status GameStatus
	turn   int
	board  Board
}

func NewGame(sz ...int) Game {
	if len(sz) != 2 {
		sz = []int{6, 7}
	}

	return Game{
		status: Initial,
		turn:   0,
		board:  NewBoard(sz[0], sz[1]),
	}
}

func (game Game) Status() GameStatus {
	return game.status
}

func (game Game) Turn() Player {
	player := None

	switch game.status {
	case Running:
		player = Player(game.turn%2 + 1)
	case Completed:
		player = Player((game.turn-1)%2 + 1)
	}

	return player
}

func (game Game) TurnCount() int {
	return game.turn
}

func (game Game) Board() Board {
	return game.board
}

func (game *Game) Start() {
	game.status = Running
	game.turn = 0
	game.board = NewBoard(game.board.RowCount(), game.board.ColCount())
}

func (game *Game) PlayTurn(column int) {
	if game.status != Running {
		return
	}

	if !game.board.inBounds(0, column) {
		return
	}

	for row := game.board.RowCount() - 1; row >= 0; row-- {
		canPlace := game.board.Get(row, column) == None

		if canPlace {
			game.board.Set(row, column, game.Turn())
			game.turn += 1

			if game.placementWins(row, column) {
				game.status = Completed
			} else if game.turn == game.board.RowCount()*game.board.ColCount() {
				game.status = Draw
			}
			break
		}
	}
}

func (game Game) placementWins(row, column int) bool {

	switch {
	case game.board.CountBidirection(row, column, North) >= 4:
		return true
	case game.board.CountBidirection(row, column, NorthEast) >= 4:
		return true
	case game.board.CountBidirection(row, column, East) >= 4:
		return true
	case game.board.CountBidirection(row, column, SouthEast) >= 4:
		return true
	}
	return false
}
