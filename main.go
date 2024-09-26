package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dstodev/go-four/c4"

	"github.com/dstodev/go-four/ui/topmenu"
)

func main() {
	game := c4.NewGame()
	game.Start()

	if _, err := tea.NewProgram(topmenu.New()).Run(); err != nil {
		fmt.Printf("Program failed: \n%v\n", err)
		os.Exit(1)
	}
}
