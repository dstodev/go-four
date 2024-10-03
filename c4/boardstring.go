package c4

import (
	"strconv"
	"strings"
)

func ToBoardString(game Game) string {
	history := game.History()

	var historyStrs []string

	for _, column := range history {
		historyStr := strconv.Itoa(column)
		historyStrs = append(historyStrs, historyStr)
	}

	return strings.Join(historyStrs, ",")
}

func FromBoardString(game Game, board_str string) Game {
	history := strings.Split(board_str, ",")

	var columns []int

	for _, columnStr := range history {
		column, err := strconv.Atoi(columnStr)

		if err == nil {
			columns = append(columns, column)
		}
	}

	if len(columns) > 0 {
		game.Start() // Resets the game
	}

	for _, column := range columns {
		game.PlayTurn(column)
	}

	return game
}
