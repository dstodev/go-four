package c4game

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/ui/optionsmenu"
)

type Outputs struct {
	Back bool
}

type button int

const (
	Place button = iota
	Back
	Quit
)

func (b button) String() string {
	return [...]string{
		"Place Token",
		"Back",
		"Quit",
	}[b]
}

type Model struct {
	outputs *Outputs
	options optionsmenu.Outputs

	cursor int
	column int

	buttons []button

	game c4.Game

	keys KeyMap
	help help.Model
}

func New(outputs *Outputs, options optionsmenu.Outputs) Model {
	game := c4.NewGame(options.Rows, options.Columns)
	game.Start()

	help := help.New()
	help.ShowAll = true

	return Model{
		outputs: outputs,
		options: options,

		cursor: 0,
		column: 0,

		buttons: []button{
			Place,
			Back,
			Quit,
		},

		game: game,

		keys: Keys,
		help: help,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Back):
			m.outputs.Back = true

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keys.Down):
			if m.cursor < len(m.buttons)-1 {
				m.cursor++
			}

		case key.Matches(msg, m.keys.Left):
			if m.column > 0 {
				m.column--
			}

		case key.Matches(msg, m.keys.Right):
			if m.column < m.game.Board().ColCount()-1 {
				m.column++
			}

		case key.Matches(msg, m.keys.Select):
			switch m.buttons[m.cursor] {
			case Place:
				m.game.PlayTurn(m.column)
				if m.game.Status() == c4.Completed || m.game.Status() == c4.Draw {
					m.buttons = []button{Back, Quit}
					m.cursor = 0
				}
			case Back:
				m.outputs.Back = true
			case Quit:
				return m, tea.Quit
			}
		}
	}

	if m.buttons[m.cursor] == Place && m.game.Status() == c4.Running {
		m.keys.Left.SetEnabled(true)
		m.keys.Right.SetEnabled(true)
	} else {
		m.keys.Left.SetEnabled(false)
		m.keys.Right.SetEnabled(false)
	}

	return m, nil
}

func (m Model) View() string {
	view := "\n"

	playerName := ""
	var playerStyle lipgloss.Style

	player1Style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.options.Player1Color))
	player2Style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.options.Player2Color))

	if m.game.Turn() == c4.One {
		playerName = m.options.Player1Name
		playerStyle = player1Style
	} else {
		playerName = m.options.Player2Name
		playerStyle = player2Style
	}

	playerName = playerStyle.Render(playerName)

	if m.game.Status() == c4.Running {
		view += fmt.Sprintf("Turn: %s (#%d)", playerName, m.game.TurnCount())
	} else if m.game.Status() == c4.Completed {
		view += fmt.Sprintf("Game over! %s wins! (Turn #%d)", playerName, m.game.TurnCount())
	} else {
		view += "Draw!"
	}

	view += "\n\n"

	view += " " + strings.Repeat("    ", m.column)
	view += " " + playerStyle.Render("↓") + " \n"

	for i, r := range m.game.Board().Rows() {
		view += "├"

		for j, c := range r {
			aboveToken := m.game.Board().Get(i, j) == c4.None &&
				m.game.Board().Get(i+1, j) != c4.None &&
				j == m.column

			bottomRowColumn := i == m.game.Board().RowCount()-1 &&
				m.game.Board().Get(i, j) == c4.None &&
				j == m.column

			placementIndicator := (aboveToken || bottomRowColumn) && m.game.Status() == c4.Running

			if placementIndicator {
				view += playerStyle.Render("↓")
			} else {
				view += " "
			}

			switch c {
			case c4.None:
				view += c.Short().String()
			case c4.One:
				view += player1Style.Render(m.options.Player1Indicator)
			case c4.Two:
				view += player2Style.Render(m.options.Player2Indicator)
			}

			if placementIndicator {
				view += playerStyle.Render("↓")
			} else {
				view += " "
			}

			if j < m.game.Board().ColCount()-1 {
				view += "│"
			} else {
				view += "┤"
			}
		}

		if i < m.game.Board().RowCount()-1 {
			view += "\n├" + strings.Repeat("───┼", m.game.Board().ColCount()-1) + "───┤"
		}

		view += "\n"
	}
	view += "├" + strings.Repeat("───┴", m.game.Board().ColCount()-1) + "───┤\n"

	view += "\n"

	for _, b := range m.buttons {
		cursor := " "

		if m.buttons[m.cursor] == b {
			cursor = ">"
		}

		view += fmt.Sprintf(" %s %s\n", cursor, b)
	}

	helpView := m.help.View(m.keys)
	// height := m.height - strings.Count(view, "\n") - strings.Count(helpView, "\n")

	// view += strings.Repeat("\n", height)
	view += "\n" + helpView

	return view
}
