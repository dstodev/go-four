package c4

type Board struct {
	rows    int
	columns int
	data    [][]Player
}

func NewBoard(rows, columns int) Board {
	data := make([][]Player, rows)

	for i := range data {
		data[i] = make([]Player, columns)
	}

	return Board{
		rows:    rows,
		columns: columns,
		data:    data,
	}
}

func (b Board) Rows() int {
	return b.rows
}

func (b Board) Columns() int {
	return b.columns
}

func (b Board) Get(row, column int) Player {
	player := None

	if b.inBounds(row, column) {
		player = b.data[row][column]
	}

	return player
}

func (b *Board) Set(row, column int, player Player) *Board {
	if b.inBounds(row, column) {
		b.data[row][column] = player
	}
	return b
}

func (b Board) inBounds(row, column int) bool {
	return row >= 0 && row < b.rows && column >= 0 && column < b.columns
}

func (b Board) IsEqual(other Board) bool {
	if b.rows != other.rows || b.columns != other.columns {
		return false
	}

	for row := 0; row < b.rows; row++ {
		for column := 0; column < b.columns; column++ {
			if b.data[row][column] != other.data[row][column] {
				return false
			}
		}
	}

	return true
}

func (b Board) Clone() Board {
	clone := NewBoard(b.rows, b.columns)

	for row := 0; row < b.rows; row++ {
		for column := 0; column < b.columns; column++ {
			clone.Set(row, column, b.Get(row, column))
		}
	}

	return clone
}

func (b Board) Neighbor(row, column int, direction Direction) Player {
	point := Point{int(row), int(column)}.Step(direction)
	return b.Get(point.Get())
}

func (b Board) CountDirection(row, column int, direction Direction) int {
	count := 0

	point := Point{row, column}

	player := b.Get(point.Get())

	for ; b.inBounds(point.Get()); point = point.Step(direction) {
		if b.Get(point.Get()) == player {
			count += 1
		} else {
			break
		}
	}

	return count
}

func (b Board) CountBidirection(row, column int, direction Direction) int {
	count := 0

	if b.inBounds(row, column) {
		count += b.CountDirection(row, column, direction)
		count += b.CountDirection(row, column, direction.Negate())
		count -= 1
	}

	return count
}
