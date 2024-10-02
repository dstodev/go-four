package c4

type Game struct {
	status   GameStatus
	turn     int
	history  []int
	board    Board
	toWin    int
	maxTurns int
}

func NewGame(sz ...int) Game {
	rows_cols := 2
	rows_cols_toWin := 3
	rows_cols_toWin_maxTurns := 4

	switch len(sz) {
	case rows_cols:
		sz = []int{sz[0], sz[1], 4, 0}

	case rows_cols_toWin:
		sz = []int{sz[0], sz[1], sz[2], 0}

	case rows_cols_toWin_maxTurns:
		// use all values (skip default case)

	default:
		sz = []int{6, 7, 4, 0}
	}

	return Game{
		status:   Initial,
		turn:     0,
		history:  []int{},
		board:    NewBoard(sz[0], sz[1]),
		toWin:    sz[2],
		maxTurns: sz[3],
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

func (game Game) History() []int {
	return game.history
}

func (game Game) TurnCount() int {
	return game.turn
}

func (game Game) ToWin() int {
	return game.toWin
}

func (game Game) MaxTurns() int {
	return game.maxTurns
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

	if !game.board.InBounds(0, column) {
		return
	}

	for row := game.board.RowCount() - 1; row >= 0; row-- {
		canPlace := game.board.Get(row, column) == None

		if canPlace {
			game.history = append(game.history, column)

			game.board.Set(row, column, game.Turn())
			game.turn += 1

			if game.placementWins(row, column) {
				game.status = Completed
			} else if game.turn == game.board.RowCount()*game.board.ColCount() ||
				game.turn == game.maxTurns {
				game.status = Draw
			}
			break
		}
	}
}

func (game Game) placementWins(row, column int) bool {
	toWin := game.toWin

	switch {
	case game.board.CountBidirection(row, column, North) >= toWin:
		return true
	case game.board.CountBidirection(row, column, NorthEast) >= toWin:
		return true
	case game.board.CountBidirection(row, column, East) >= toWin:
		return true
	case game.board.CountBidirection(row, column, SouthEast) >= toWin:
		return true
	}
	return false
}
