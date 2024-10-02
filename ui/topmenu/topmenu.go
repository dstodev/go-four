package topmenu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/dstodev/go-four/ui"
	"github.com/dstodev/go-four/ui/c4game"
	"github.com/dstodev/go-four/ui/optionsmenu"
	"github.com/dstodev/go-four/util"
)

type button int

const (
	buttonNewGame button = iota
	buttonOptions
	buttonHelp
	buttonQuit
)

func (b button) String() string {
	return [...]string{
		"New Game",
		"Options",
		"Help",
		"Quit",
	}[b]
}

type menu int

const (
	menuMain menu = iota
	menuGame
	menuOptions
)

type Model struct {
	cursor      *util.Cursor
	currentMenu menu

	buttons []button

	maxHeight int

	game *c4game.Model

	options    optionsmenu.Model
	optionsOut *optionsmenu.Outputs

	keys ui.KeyMap
	help help.Model
}

func New() Model {
	optionsOut := &optionsmenu.Outputs{}

	rows, _ := util.ViewportSize()

	help := help.New()
	help.ShowAll = true

	keys := ui.DefaultKeys
	keys.Left.SetEnabled(false)
	keys.Right.SetEnabled(false)
	keys.Back.SetEnabled(false)
	keys.Quit.SetKeys("esc", "ctrl+c")
	keys.Quit.SetHelp("esc/ctrl+c", "Quit")

	return Model{
		cursor:      util.NewCursor(),
		currentMenu: menuMain,

		buttons: []button{
			buttonNewGame,
			buttonOptions,
			buttonHelp,
			buttonQuit,
		},

		maxHeight: rows,

		game: nil,

		options:    optionsmenu.New(optionsOut, rows),
		optionsOut: optionsOut,

		keys: keys,
		help: help,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.maxHeight = msg.Height

		m.options, _ = m.options.Update(msg)
		if m.game != nil {
			*m.game, _ = m.game.Update(msg)
		}
		return m, nil

	case ui.SetFullHelpMsg:
		show := bool(msg)
		m.help.ShowAll = show

		m.options, _ = m.options.Update(msg)
		if m.game != nil {
			*m.game, _ = m.game.Update(msg)
		}
		return m, nil

	case ui.BackMsg:
		m.currentMenu = menuMain
		return m, nil
	}

	switch m.currentMenu {
	case menuMain:
		cmd = m.internalUpdate(msg)

	case menuGame:
		*m.game, cmd = m.game.Update(msg)

	case menuOptions:
		m.options, cmd = m.options.Update(msg)
	}

	return m, cmd
}

func (m *Model) internalUpdate(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return tea.Quit

		case key.Matches(msg, m.keys.Help):
			cmd = ui.SetFullHelpCmd(!m.help.ShowAll)

		case key.Matches(msg, m.keys.Up):
			m.cursor.MoveUp()

		case key.Matches(msg, m.keys.Down):
			m.cursor.MoveDown()

		case key.Matches(msg, m.keys.Select):
			switch m.buttons[m.cursor.Row()] {
			case buttonNewGame:
				m.game = c4game.New(*m.optionsOut, m.maxHeight)
				cmd = ui.SetFullHelpCmd(m.help.ShowAll)
				m.currentMenu = menuGame

			case buttonOptions:
				m.currentMenu = menuOptions

			case buttonHelp:
				cmd = ui.SetFullHelpCmd(!m.help.ShowAll)

			case buttonQuit:
				return tea.Quit
			}
		}
	}

	m.cursor.ConstrainRow(0, len(m.buttons))

	return cmd
}

func (m Model) View() string {
	switch m.currentMenu {
	case menuMain:
		return m.internalView()

	case menuGame:
		return m.game.View()

	case menuOptions:
		return m.options.View()
	}

	return ""
}

func (m Model) internalView() string {
	view := "\n Go Four!\n\n"

	for _, b := range m.buttons {
		cursor := " "

		if m.buttons[m.cursor.Row()] == b {
			cursor = ">"
		}

		view += fmt.Sprintf(" %s %s\n", cursor, b)
	}

	helpView := m.help.View(m.keys)
	view += "\n" + helpView

	return strings.Join(util.LastNLines(view, m.maxHeight), "\n")
}
