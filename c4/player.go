package c4

type Player uint8

const (
	None Player = iota
	One
	Two
)

func (player Player) String() string {
	return [...]string{
		"None",
		"One",
		"Two",
	}[player]
}

func (player Player) Negate() Player {
	return [...]Player{
		None,
		Two,
		One,
	}[player]
}

type PlayerShort uint8

func (player Player) Short() PlayerShort {
	return PlayerShort(player)
}

func (player PlayerShort) String() string {
	return [...]string{
		" ",
		"A",
		"B",
	}[player]
}
