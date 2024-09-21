package optionsmenu

type action int

const (
	Back action = iota

	EnterRows
	EnterColumns

	EnterPlayer1Name
	EnterPlayer1Indicator
	EnterPlayer1Color

	EnterPlayer2Name
	EnterPlayer2Indicator
	EnterPlayer2Color
)

func (b action) String() string {
	return [...]string{
		"Back",

		"Rows",
		"Columns",

		"Player 1 name",
		"Player 1 indicator",
		"Player 1 color",

		"Player 2 name",
		"Player 2 indicator",
		"Player 2 color",
	}[b]
}

func (b action) opposite() action {
	switch b {
	case EnterPlayer1Name:
		return EnterPlayer2Name
	case EnterPlayer1Indicator:
		return EnterPlayer2Indicator
	case EnterPlayer1Color:
		return EnterPlayer2Color

	case EnterPlayer2Name:
		return EnterPlayer1Name
	case EnterPlayer2Indicator:
		return EnterPlayer1Indicator
	case EnterPlayer2Color:
		return EnterPlayer1Color

	default:
		return -1
	}
}
