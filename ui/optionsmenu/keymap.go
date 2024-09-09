package optionsmenu

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Select key.Binding
	Quit   key.Binding
	Help   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right, k.Select},
		{k.Quit, k.Help},
	}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "w", "k"),
		key.WithHelp("↑/w/k", "Move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "s", "j"),
		key.WithHelp("↓/s/j", "Move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "a", "h"),
		key.WithHelp("←/a/h", "Move left"),
		key.WithDisabled(),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "d", "l"),
		key.WithHelp("→/d/l", "Move right"),
		key.WithDisabled(),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("space/enter", "Select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc/C-c", "Quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?", "h"),
		key.WithHelp("?/h", "Help"),
	),
}
