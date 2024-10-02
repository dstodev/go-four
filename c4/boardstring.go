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

func FromBoardString(game Game, board_str string) Game {
	if board_str == "" {
		return game
	}

	game.Start() // Resets the game

	history := strings.Split(board_str, ",")

	for _, column := range history {
		column, err := strconv.Atoi(column)
		if err == nil {
			game.PlayTurn(column)
		}
	}

	return game
}
