package optionsmenu

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dstodev/go-four/ui"
	tb "github.com/dstodev/go-four/ui/textbox"
	"github.com/dstodev/go-four/util"
)

type Outputs struct {
	Rows     int
	Columns  int
	ToWin    int
	MaxTurns int

	Player1Name      string
	Player1Indicator string
	Player1Color     string

	Player2Name      string
	Player2Indicator string
	Player2Color     string
}

type Model struct {
	outputs *Outputs

	cursor *util.Cursor

	buttons []action

	inputs         map[action]tb.Model
	currentTextbox action

	maxHeight int

	keys ui.KeyMap
	help help.Model
}

func New(outputs *Outputs, height int) Model {
	*outputs = Outputs{
		Rows:     6,
		Columns:  7,
		ToWin:    4,
		MaxTurns: 0,

		Player1Name:      "Player One",
		Player1Indicator: "A",
		Player1Color:     "009fff",

		Player2Name:      "Player Two",
		Player2Indicator: "B",
		Player2Color:     "ff9400",
	}

	help := help.New()
	help.ShowAll = true

	inputs := map[action]tb.Model{
		EnterRows:     tb.NewInteger(&outputs.Rows, 1, tb.ConstrainGreaterEq(4)).WithLabel(EnterRows.String() + "    (4-9): "),
		EnterColumns:  tb.NewInteger(&outputs.Columns, 1, tb.ConstrainGreaterEq(4)).WithLabel(EnterColumns.String() + " (4-9): "),
		EnterToWin:    tb.NewInteger(&outputs.ToWin, 1, tb.ConstrainGreaterEq(3)).WithLabel(EnterToWin.String() + "  (3-9): "),
		EnterMaxTurns: tb.NewInteger(&outputs.MaxTurns, 3, tb.ConstrainGreaterEqZero, tb.ConstrainLessEq(100)).WithLabel(EnterMaxTurns.String() + " (0-100): "),

		EnterPlayer1Name:      tb.NewString(&outputs.Player1Name, 20),
		EnterPlayer1Indicator: tb.NewString(&outputs.Player1Indicator, 1),
		EnterPlayer1Color:     tb.NewColor(&outputs.Player1Color),

		EnterPlayer2Name:      tb.NewString(&outputs.Player2Name, 20),
		EnterPlayer2Indicator: tb.NewString(&outputs.Player2Indicator, 1),
		EnterPlayer2Color:     tb.NewColor(&outputs.Player2Color),
	}

	buttons := []action{
		Back,
	}

	for b := range inputs {
		buttons = append(buttons, b)
	}
	sort.Slice(buttons, func(i, j int) bool {
		return buttons[i] < buttons[j]
	})

	keys := ui.DefaultKeys
	resetKeymap(&keys)

	return Model{
		outputs: outputs,

		cursor: util.NewCursor(),

		buttons: buttons,

		inputs:         inputs,
		currentTextbox: -1,

		maxHeight: height,

		keys: keys,
		help: help,
	}
}

func resetKeymap(km *ui.KeyMap) {
	*km = ui.DefaultKeys
	km.Left.SetEnabled(false)
	km.Right.SetEnabled(false)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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
			m.cursor.MoveUp()

		case key.Matches(msg, m.keys.Down):
			m.cursor.MoveDown()

		case key.Matches(msg, m.keys.Select):
			b := m.buttons[m.cursor.Row()]

			switch b {
			case Back:
				cmd = ui.BackCmd

			default:
				cmd = m.toggleInput(b)
				return m, cmd
			}
		}
	}

	m.cursor.ConstrainRow(0, len(m.buttons))

	if m.anyInputFocused() {
		box := m.inputs[m.currentTextbox]
		m.inputs[m.currentTextbox], cmd = box.Update(msg)
	}

	return m, cmd
}

func (m *Model) toggleInput(b action) tea.Cmd {
	var cmd tea.Cmd

	box := m.inputs[b]

	if m.anyInputFocused() {
		m.currentTextbox = -1
		resetKeymap(&m.keys)

		box.Leave()

		if opposite := b.Opposite(); opposite != -1 {
			oppositeBox := m.inputs[opposite]
			box.AssertDifferent(oppositeBox)
		}
	} else {
		m.currentTextbox = b
		m.constrainKeymap()

		cmd = box.Enter()
	}

	m.inputs[b] = box

	return cmd
}

func (m Model) anyInputFocused() bool {
	return m.currentTextbox != -1
}

func (m *Model) constrainKeymap() {
	m.keys.Back.SetEnabled(false)
	m.keys.Up.SetEnabled(false)
	m.keys.Down.SetEnabled(false)
	m.keys.Select.SetKeys("esc", "enter")
	m.keys.Select.SetHelp("esc/enter", "Confirm")
}

func (m Model) View() string {
	view := "\n Go Four options:\n\n"

	player1Style := lipgloss.NewStyle().Foreground(lipgloss.Color("#" + m.outputs.Player1Color))
	player2Style := lipgloss.NewStyle().Foreground(lipgloss.Color("#" + m.outputs.Player2Color))

	for _, b := range m.buttons {
		cursor := " "

		if m.buttons[m.cursor.Row()] == b {
			cursor = ">"
		}

		box := m.inputs[b]

		switch b {
		case Back:
			view += fmt.Sprintf(" %s %s\n\n", cursor, b) // extra newline

		case EnterMaxTurns:
			view += fmt.Sprintf(" %s %s\n\n", cursor, box.View()) // extra newline

		case EnterPlayer1Name:
			box.SetLabel(player1Style.Render(EnterPlayer1Name.String() + ": "))
			view += fmt.Sprintf(" %s %s\n", cursor, box.View())

		case EnterPlayer1Indicator:
			box.SetLabel(player1Style.Render(EnterPlayer1Indicator.String() + ": "))
			view += fmt.Sprintf(" %s %s\n", cursor, box.View())

		case EnterPlayer1Color:
			box.SetLabel(player1Style.Render(EnterPlayer1Color.String()+": ") + "#")
			view += fmt.Sprintf(" %s %s\n\n", cursor, box.View()) // extra newline

		case EnterPlayer2Name:
			box.SetLabel(player2Style.Render(EnterPlayer2Name.String() + ": "))
			view += fmt.Sprintf(" %s %s\n", cursor, box.View())

		case EnterPlayer2Indicator:
			box.SetLabel(player2Style.Render(EnterPlayer2Indicator.String() + ": "))
			view += fmt.Sprintf(" %s %s\n", cursor, box.View())

		case EnterPlayer2Color:
			box.SetLabel(player2Style.Render(EnterPlayer2Color.String()+": ") + "#")
			view += fmt.Sprintf(" %s %s\n", cursor, box.View())

		default:
			view += fmt.Sprintf(" %s %s\n", cursor, box.View())
		}
	}

	helpView := m.help.View(m.keys)
	view += "\n" + helpView

	return strings.Join(util.LastNLines(view, m.maxHeight), "\n")
}
