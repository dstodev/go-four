package c4

type Direction int

const (
	NoDirection Direction = iota
	North
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

func (status Direction) String() string {
	return [...]string{
		"None",
		"North",
		"NorthEast",
		"East",
		"SouthEast",
		"South",
		"SouthWest",
		"West",
		"NorthWest",
	}[status]
}

func (direction Direction) Negate() Direction {
	return [...]Direction{
		NoDirection,
		South,
		SouthWest,
		West,
		NorthWest,
		North,
		NorthEast,
		East,
		SouthEast,
	}[direction]
}

func (direction Direction) OffsetRow() int {
	return [...]int{
		0,
		-1,
		-1,
		0,
		1,
		1,
		1,
		0,
		-1,
	}[direction]
}

func (direction Direction) OffsetColumn() int {
	return [...]int{
		0,
		0,
		1,
		1,
		1,
		0,
		-1,
		-1,
		-1,
	}[direction]
}

func (direction Direction) Offset() Point {
	return Point{direction.OffsetRow(), direction.OffsetColumn()}
}
