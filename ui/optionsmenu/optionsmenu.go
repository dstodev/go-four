package optionsmenu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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

type Model struct {
	outputs *Outputs

	cursor int
	height int

	buttons []action

	inputs         map[action]textbox
	currentTextbox action

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
		Player1Color:     "009fff",

		Player2Name:      "Player Two",
		Player2Indicator: "B",
		Player2Color:     "ff7f00",
	}

	help := help.New()
	help.ShowAll = true

	inputs := map[action]textbox{
		EnterRows:    newIntegerInput(&outputs.Rows),
		EnterColumns: newIntegerInput(&outputs.Columns),

		EnterPlayer1Name:      newStringInput(&outputs.Player1Name, 10),
		EnterPlayer1Indicator: newStringInput(&outputs.Player1Indicator, 1),
		EnterPlayer1Color:     newColorInput(&outputs.Player1Color),

		EnterPlayer2Name:      newStringInput(&outputs.Player2Name, 10),
		EnterPlayer2Indicator: newStringInput(&outputs.Player2Indicator, 1),
		EnterPlayer2Color:     newColorInput(&outputs.Player2Color),
	}

	buttons := []action{
		Back,
	}

	for b := range inputs {
		buttons = append(buttons, b)
	}

	return Model{
		outputs: outputs,

		cursor: 0,
		height: 18,

		buttons: []action{
			Back,
		},

		inputs:         inputs,
		currentTextbox: -1,

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
			textbox := m.inputs[b]

			switch b {
			case Back:
				m.outputs.Back = true

			default:
				if m.currentTextbox == -1 {
					m.currentTextbox = b
					m.enterTextbox()
					textbox.enter()
					return m, nil
				} else {
					m.currentTextbox = -1
					m.leaveTextbox()
					textbox.leave()
				}
			}
		}
	}

	var cmd tea.Cmd

	if m.currentTextbox != -1 {
		textbox := m.inputs[m.currentTextbox]
		cmd = textbox.updateModel(msg)
	}

	return m, cmd
}

func (m *Model) enterTextbox() {
	// Constrain keymap (would conflict with text input)
	m.keys.Back.SetEnabled(false)
	m.keys.Up.SetEnabled(false)
	m.keys.Down.SetEnabled(false)
	m.keys.Select.SetKeys("esc", "enter")
	m.keys.Select.SetHelp("esc/enter", "Confirm")
}

func (m *Model) leaveTextbox() {
	// Reset keymap
	m.keys = Keys
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
