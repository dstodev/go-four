package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct{}

func BackCmd() tea.Msg {
	return BackMsg{}
}

type SetFullHelpMsg bool

func SetFullHelpCmd(b bool) tea.Cmd {
	return func() tea.Msg { return SetFullHelpMsg(b) }
}
