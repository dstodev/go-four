package c4game

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/ui"
	"github.com/dstodev/go-four/ui/optionsmenu"
	"github.com/dstodev/go-four/util"
)

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
	options optionsmenu.Outputs

	cursor int
	column int

	buttons []button

	maxHeight int

	game c4.Game

	keys KeyMap
	help help.Model
}

func New(options optionsmenu.Outputs, height int) Model {
	game := c4.NewGame(options.Rows, options.Columns, options.ToWin, options.MaxTurns)
	game.Start()

	help := help.New()
	help.ShowAll = true

	return Model{
		options: options,

		cursor: 0,
		column: 0,

		buttons: []button{
			Place,
			Back,
			Quit,
		},

		maxHeight: height,

		game: game,

		keys: Keys,
		help: help,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case ui.SetFullHelpMsg:
		m.help.ShowAll = bool(msg)

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.maxHeight = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Back):
			cmd = ui.BackCmd

		case key.Matches(msg, m.keys.Help):
			cmd = ui.SetFullHelpCmd(!m.help.ShowAll)

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
				cmd = ui.BackCmd

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

	return m, cmd
}

func (m Model) View() string {
	view := "\n"

	playerName := ""
	var playerStyle lipgloss.Style

	player1Style := lipgloss.NewStyle().Foreground(lipgloss.Color("#" + m.options.Player1Color))
	player2Style := lipgloss.NewStyle().Foreground(lipgloss.Color("#" + m.options.Player2Color))

	switch m.game.Turn() {
	case c4.None:
		// This style is used when there is a draw
		playerStyle = lipgloss.NewStyle()
	case c4.One:
		playerName = m.options.Player1Name
		playerStyle = player1Style
	case c4.Two:
		playerName = m.options.Player2Name
		playerStyle = player2Style
	}

	playerName = playerStyle.Render(playerName)

	board := m.game.Board()

	view += " "

	var partOfWin [][]bool

	switch m.game.Status() {
	case c4.Running:
		var maxTurns string

		if m.game.MaxTurns() > 0 {
			maxTurns = fmt.Sprintf(" of %d", m.game.MaxTurns())
		}

		view += fmt.Sprintf("Turn: %s (#%d%s)", playerName, m.game.TurnCount(), maxTurns)

	case c4.Completed:
		view += fmt.Sprintf("Game over! %s wins! (Turn #%d)", playerName, m.game.TurnCount())

		// Get last placed token
		winPoint := c4.NewPoint(0, m.column)

		for ; winPoint.Row < board.RowCount(); winPoint = winPoint.Step(c4.South) {
			if board.Get(winPoint.Get()) == m.game.Turn() {
				break
			}
		}

		// Get winning chain(s)
		partOfWin = make([][]bool, board.RowCount())

		for i := range partOfWin {
			partOfWin[i] = make([]bool, board.ColCount())
		}

		for direction := c4.North; direction < c4.NorthWest; direction++ {
			if board.CountBidirection(winPoint.Row, winPoint.Col, direction) >= m.game.ToWin() {
				for crawl := winPoint; board.Get(crawl.Get()) == m.game.Turn(); crawl = crawl.Step(direction) {
					partOfWin[crawl.Row][crawl.Col] = true
				}
			}
		}

	default:
		view += fmt.Sprintf("Draw on turn %d!", m.game.TurnCount())
	}

	view += "\n\n"

	boardLeftPadLength := 4
	boardLeftPad := strings.Repeat(" ", boardLeftPadLength)

	view += " " + strings.Repeat("    ", m.column)
	view += boardLeftPad + " " + playerStyle.Render("↓") + " \n"

	for i, r := range board.Rows() {
		view += boardLeftPad + "┤"

		for j, c := range r {
			aboveToken := board.Get(i, j) == c4.None &&
				board.Neighbor(i, j, c4.South) != c4.None &&
				j == m.column

			bottomRowColumn := i == board.RowCount()-1 &&
				board.Get(i, j) == c4.None &&
				j == m.column

			placementIndicator := (aboveToken || bottomRowColumn) && m.game.Status() == c4.Running

			if placementIndicator {
				view += playerStyle.Render("↓")
			} else {
				view += " "
			}

			if len(partOfWin) > 0 && partOfWin[i][j] {
				player1Style = player1Style.Bold(true).Underline(true)
				player2Style = player2Style.Bold(true).Underline(true)
			} else {
				player1Style = player1Style.Bold(false).Underline(false)
				player2Style = player2Style.Bold(false).Underline(false)
			}

			switch c {
			case c4.None:
				view += " "
			// Must render using fixed styles here to preserve players' already-placed token colors
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

			if j < board.ColCount()-1 {
				view += "│"
			} else {
				view += "├"
			}
		}

		if i < board.RowCount()-1 {
			view += "\n" + boardLeftPad + "├" + strings.Repeat("───┼", board.ColCount()-1) + "───┤"
		}

		view += "\n"
	}
	view += boardLeftPad + "├" + strings.Repeat("───┴", board.ColCount()-1) + "───┤\n"

	view += "\n"

	for _, b := range m.buttons {
		cursor := " "

		if m.buttons[m.cursor] == b {
			cursor = ">"
		}

		view += fmt.Sprintf(" %s %s\n", cursor, b)
	}

	helpView := m.help.View(m.keys)
	view += "\n" + helpView

	height := util.CountLines(view)
	height = util.Min(m.maxHeight, height)
	height = util.Max(0, height)

	return util.LastNLines(view, height)
}
