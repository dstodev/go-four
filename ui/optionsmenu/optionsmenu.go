package optionsmenu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Outputs struct {
	Back bool

	Rows    int
	Columns int

	Player1Name      string
	Player1Indicator string
	Player1Color     string

	Player2Name      string
	Player2Indicator string
	Player2Color     string
}

type button int

const (
	Back button = iota

	EnterRows
	EnterColumns

	EnterPlayer1Name
	EnterPlayer1Indicator
	EnterPlayer1Color

	EnterPlayer2Name
	EnterPlayer2Indicator
	EnterPlayer2Color
)

func (b button) String() string {
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

type Model struct {
	outputs *Outputs

	cursor int

	buttons []button
	height  int

	currentTextbox button

	rows    textinput.Model
	columns textinput.Model

	player1Name      textinput.Model
	player1Indicator textinput.Model
	player1Color     textinput.Model

	player2Name      textinput.Model
	player2Indicator textinput.Model
	player2Color     textinput.Model

	keys KeyMap
	help help.Model
}

func New(outputs *Outputs) Model {
	*outputs = Outputs{
		Back: false,

		Rows:    6,
		Columns: 7,

		Player1Name:      "Player One",
		Player1Indicator: "A",
		Player1Color:     "ff0000",

		Player2Name:      "Player Two",
		Player2Indicator: "B",
		Player2Color:     "00ff00",
	}

	help := help.New()
	help.ShowAll = true

	rows := textinput.New()
	rows.Placeholder = strconv.Itoa(outputs.Rows)
	rows.CharLimit = 1
	rows.Width = 1
	rows.Prompt = "   " + EnterRows.String() + " (1-9): "

	columns := textinput.New()
	columns.Placeholder = strconv.Itoa(outputs.Columns)
	columns.CharLimit = 1
	columns.Width = 1
	columns.Prompt = EnterColumns.String() + " (1-9): "

	player1Name := textinput.New()
	player1Name.Placeholder = outputs.Player1Name
	player1Name.CharLimit = 10
	player1Name.Width = 10

	player1Indicator := textinput.New()
	player1Indicator.Placeholder = outputs.Player1Indicator
	player1Indicator.CharLimit = 1
	player1Indicator.Width = 1

	player1Color := textinput.New()
	player1Color.Placeholder = outputs.Player1Color
	player1Color.CharLimit = 7
	player1Color.Width = 7

	player2Name := textinput.New()
	player2Name.Placeholder = outputs.Player2Name
	player2Name.CharLimit = 10
	player2Name.Width = 10

	player2Indicator := textinput.New()
	player2Indicator.Placeholder = outputs.Player2Indicator
	player2Indicator.CharLimit = 1
	player2Indicator.Width = 1

	player2Color := textinput.New()
	player2Color.Placeholder = outputs.Player2Color
	player2Color.CharLimit = 7
	player2Color.Width = 7

	return Model{
		outputs: outputs,

		cursor: 0,

		buttons: []button{
			Back,

			EnterRows,
			EnterColumns,

			EnterPlayer1Name,
			EnterPlayer1Indicator,
			EnterPlayer1Color,

			EnterPlayer2Name,
			EnterPlayer2Indicator,
			EnterPlayer2Color,
		},
		height: 18,

		currentTextbox: -1,

		rows:    rows,
		columns: columns,

		player1Name:      player1Name,
		player1Indicator: player1Indicator,
		player1Color:     player1Color,

		player2Name:      player2Name,
		player2Indicator: player2Indicator,
		player2Color:     player2Color,

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

		case key.Matches(msg, m.keys.Select):
			b := m.buttons[m.cursor]

			switch b {
			case Back:
				m.outputs.Back = true

			default:
				if m.currentTextbox == -1 {
					m.currentTextbox = b
					m.enterTextbox(m.toTextbox(b))
					return m, nil
				} else {
					m.currentTextbox = -1
					m.leaveTextbox(m.toTextbox(b))
				}
			}
		}
	}

	var cmd tea.Cmd

	if m.currentTextbox != -1 {
		textbox := m.toTextbox(m.currentTextbox)
		*textbox, cmd = textbox.Update(msg)
	}

	return m, cmd
}

func (m *Model) enterTextbox(textbox *textinput.Model) {
	// Constrain keymap (would conflict with text input)
	m.keys.Back.SetEnabled(false)
	m.keys.Up.SetEnabled(false)
	m.keys.Down.SetEnabled(false)
	m.keys.Select.SetKeys("esc", "enter")
	m.keys.Select.SetHelp("esc/enter", "Confirm")

	textbox.Reset()
	textbox.Focus()
	textbox.CursorEnd()
}

func (m *Model) leaveTextbox(textbox *textinput.Model) {
	var opposite *textinput.Model

	switch textbox {
	case &m.player1Name:
		opposite = &m.player2Name
	case &m.player1Indicator:
		opposite = &m.player2Indicator
	case &m.player1Color:
		opposite = &m.player2Color
	case &m.player2Name:
		opposite = &m.player1Name
	case &m.player2Indicator:
		opposite = &m.player1Indicator
	case &m.player2Color:
		opposite = &m.player1Color
	}

	if opposite != nil && strings.EqualFold(valueOrPlaceholder(textbox), valueOrPlaceholder(opposite)) {
		textbox.Reset()

		// If the reset causes a conflict, use the placeholder of the opposite textbox
		if strings.EqualFold(valueOrPlaceholder(textbox), valueOrPlaceholder(opposite)) {
			textbox.SetValue(opposite.Placeholder)
		}
	}

	m.updateOption(textbox)

	// Reset keymap
	m.keys = Keys

	textbox.Blur()
}

func (m *Model) updateOption(textbox *textinput.Model) {
	switch textbox {
	case &m.rows:
		if rows, err := strconv.Atoi(valueOrPlaceholder(textbox)); err == nil && rows > 0 {
			m.outputs.Rows = rows
		} else {
			textbox.Reset()
		}

	case &m.columns:
		if columns, err := strconv.Atoi(valueOrPlaceholder(textbox)); err == nil && columns > 0 {
			m.outputs.Columns = columns
		} else {
			textbox.Reset()
		}

	case &m.player1Name:
		m.outputs.Player1Name = valueOrPlaceholder(textbox)

	case &m.player1Indicator:
		m.outputs.Player1Indicator = valueOrPlaceholder(textbox)

	case &m.player1Color:
		value := valueOrPlaceholder(textbox)
		regexp.Compile("[0-9a-fA-F]{6}")

		if !regexp.MustCompile("[0-9a-fA-F]{6}").MatchString(value) {
			textbox.Reset()
		}

		m.outputs.Player1Color = valueOrPlaceholder(textbox)

	case &m.player2Name:
		m.outputs.Player2Name = valueOrPlaceholder(textbox)

	case &m.player2Indicator:
		m.outputs.Player2Indicator = valueOrPlaceholder(textbox)

	case &m.player2Color:
		value := valueOrPlaceholder(textbox)
		regexp.Compile("[0-9a-fA-F]{6}")

		if !regexp.MustCompile("[0-9a-fA-F]{6}").MatchString(value) {
			textbox.Reset()
		}

		m.outputs.Player2Color = valueOrPlaceholder(textbox)
	}
}

func valueOrPlaceholder(input *textinput.Model) string {
	if input.Value() == "" {
		return input.Placeholder
	}
	return input.Value()
}

func (m *Model) toTextbox(button button) *textinput.Model {
	switch button {
	case EnterRows:
		return &m.rows

	case EnterColumns:
		return &m.columns

	case EnterPlayer1Name:
		return &m.player1Name

	case EnterPlayer1Indicator:
		return &m.player1Indicator

	case EnterPlayer1Color:
		return &m.player1Color

	case EnterPlayer2Name:
		return &m.player2Name

	case EnterPlayer2Indicator:
		return &m.player2Indicator

	case EnterPlayer2Color:
		return &m.player2Color
	}

	return nil
}

func (m Model) View() string {
	view := "\nGo Four options:\n\n"

	player1Style := lipgloss.NewStyle().Foreground(lipgloss.Color("#" + m.outputs.Player1Color))
	player2Style := lipgloss.NewStyle().Foreground(lipgloss.Color("#" + m.outputs.Player2Color))

	for _, b := range m.buttons {
		cursor := " "

		if m.buttons[m.cursor] == b {
			cursor = ">"
		}

		switch b {
		case Back:
			view += fmt.Sprintf(" %s %s\n\n", cursor, b)

		case EnterColumns:
			textbox := m.toTextbox(b)
			view += fmt.Sprintf(" %s %s\n\n", cursor, textbox.View()) // extra newline

		case EnterPlayer1Name:
			textbox := m.toTextbox(b)
			textbox.Prompt = player1Style.Render(EnterPlayer1Name.String() + ": ")
			view += fmt.Sprintf(" %s %s\n", cursor, textbox.View())

		case EnterPlayer1Indicator:
			textbox := m.toTextbox(b)
			textbox.Prompt = player1Style.Render(EnterPlayer1Indicator.String() + ": ")
			view += fmt.Sprintf(" %s %s\n", cursor, textbox.View())

		case EnterPlayer1Color:
			textbox := m.toTextbox(b)
			textbox.Prompt = player1Style.Render(EnterPlayer1Color.String()+": ") + "#"
			view += fmt.Sprintf(" %s %s\n\n", cursor, textbox.View()) // extra newline

		case EnterPlayer2Name:
			textbox := m.toTextbox(b)
			textbox.Prompt = player2Style.Render(EnterPlayer2Name.String() + ": ")
			view += fmt.Sprintf(" %s %s\n", cursor, textbox.View())

		case EnterPlayer2Indicator:
			textbox := m.toTextbox(b)
			textbox.Prompt = player2Style.Render(EnterPlayer2Indicator.String() + ": ")
			view += fmt.Sprintf(" %s %s\n", cursor, textbox.View())

		case EnterPlayer2Color:
			textbox := m.toTextbox(b)
			textbox.Prompt = player2Style.Render(EnterPlayer2Color.String()+": ") + "#"
			view += fmt.Sprintf(" %s %s\n", cursor, textbox.View())

		default:
			textbox := m.toTextbox(b)
			view += fmt.Sprintf(" %s %s\n", cursor, textbox.View())
		}
	}

	helpView := m.help.View(m.keys)
	height := m.height - strings.Count(view, "\n") - strings.Count(helpView, "\n")

	view += strings.Repeat("\n", height)
	view += helpView

	return view
}
