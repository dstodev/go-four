package c4

type Direction int

func (status Direction) String() string {
	return [...]string{
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

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

func (direction Direction) Negate() Direction {
	return [...]Direction{
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
		1,
		1,
		1,
		0,
		-1,
		-1,
		-1,
	}[direction]
}
