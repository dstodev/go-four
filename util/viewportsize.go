package util

import (
	"os"

	"github.com/charmbracelet/x/term"
)

func ViewportSize() (rows, columns int) {
	cols, rows, err := term.GetSize(os.Stdout.Fd())

	if err != nil {
		panic(err)
	}

	return rows, cols
}
