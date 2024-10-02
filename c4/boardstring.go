package c4

import (
	"strconv"
	"strings"
)

func ToBoardString(game Game) string {
	history := game.History()

	var history_strs []string

	for _, column := range history {
		history_str := strconv.Itoa(column)
		history_strs = append(history_strs, history_str)
	}

	return strings.Join(history_strs, ",")
}

func FromBoardString(rows, columns int, board_str string) Game {
	game := NewGame(rows, columns)

	if board_str == "" {
		return game
	}

	history := strings.Split(board_str, ",")

	if len(history) > 0 {
		game.Start()
	}

	for _, column := range history {
		column, err := strconv.Atoi(column)
		if err != nil {
		}
		game.PlayTurn(column)
	}

	return game
}
