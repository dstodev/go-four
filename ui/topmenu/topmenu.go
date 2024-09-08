package topmenu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dstodev/go-four/ui/c4_game"
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
	cursor int

	buttons []button
	menu    menu
	height  int

	game c4_game.Model

	keys KeyMap
	help help.Model
}

func New() Model {
	return Model{
		cursor: 0,

		buttons: []button{
			buttonNewGame,
			buttonOptions,
			buttonHelp,
			buttonQuit,
		},
		menu:   menuMain,
		height: 10,

		game: c4_game.New(),

		keys: Keys,
		help: help.New(),
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
	}

	switch m.menu {
	case menuMain:
		return m.internalUpdate(msg)

	case menuGame:
		var model tea.Model
		model, cmd = m.game.Update(msg)

		m.game = model.(c4_game.Model)

		if m.game.Back {
			m.menu = menuMain
			m.game.Back = false
		}
	}

	return m, cmd
}

func (m *Model) internalUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keys.Down):
			if m.cursor < (len(m.buttons) - 1) {
				m.cursor++
			}

		case key.Matches(msg, m.keys.Select):
			switch m.buttons[m.cursor] {
			case buttonNewGame:
				m.game = c4_game.New()
				m.menu = menuGame
			case buttonOptions:
				// TODO: Allow changing board size, player character
			case buttonHelp:
				m.help.ShowAll = !m.help.ShowAll
			case buttonQuit:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.menu {
	case menuMain:
		return m.internalView()
	case menuGame:
		return m.game.View()
	}

	return ""
}

func (m Model) internalView() string {
	view := "\nGo Four!\n\n"

	for _, b := range m.buttons {
		cursor := " "

		if m.buttons[m.cursor] == b {
			cursor = ">"
		}

		view += fmt.Sprintf(" %s %s\n", cursor, b)
	}

	helpView := m.help.View(m.keys)
	height := m.height - strings.Count(view, "\n") - strings.Count(helpView, "\n")

	view += strings.Repeat("\n", height)
	view += helpView

	return view
}
