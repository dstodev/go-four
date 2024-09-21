package textbox

import (
	"regexp"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	input  textinput.Model
	update func(t *textinput.Model)
}

func (m Model) Init() tea.Cmd {
	//return textinput.Blink
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.input.View()
}

func (m *Model) updateOutput() {
	m.update(&m.input)
}

func enter(t *textinput.Model) {
	t.Reset()
	t.Focus()
	t.CursorEnd()
}

func leave(t *textinput.Model) {
	// opposite := textbox.opposite()

	// if opposite != nil && strings.EqualFold(valueOrPlaceholder(textbox), valueOrPlaceholder(opposite)) {
	// 	textbox.Reset()

	// 	// If the reset causes a conflict, use the placeholder of the opposite textbox
	// 	if strings.EqualFold(valueOrPlaceholder(textbox), valueOrPlaceholder(opposite)) {
	// 		textbox.SetValue(opposite.Placeholder)
	// 	}
	// }

	//m.updateOutput()

	t.Blur()
}

func newIntegerInput(output *int) Model {
	input := textinput.New()
	input.Placeholder = strconv.Itoa(*output)
	input.CharLimit = 1
	input.Width = 1
	//input.Prompt = "   " + EnterRows.String() + " (1-9): "

	return Model{
		input: input,
		update: func(t *textinput.Model) {
			if value, err := strconv.Atoi(valueOrPlaceholder(&input)); err == nil && value > 0 {
				*output = value
			} else {
				t.Reset()
			}
		},
	}
}

func newStringInput(output *string, width int) textbox {
	input := textinput.New()
	input.Placeholder = *output
	input.CharLimit = width
	input.Width = width

	return textbox{
		input: input,
		update: func(t *textbox) {
			*output = valueOrPlaceholder(&t.input)
		},
	}
}

func newColorInput(output *string) textbox {
	input := textinput.New()
	input.Placeholder = *output
	input.CharLimit = 6
	input.Width = 6

	return textbox{
		input: input,
		update: func(t *textbox) {
			value := valueOrPlaceholder(&t.input)
			regexp.Compile("[0-9a-fA-F]{6}")

			if regexp.MustCompile("[0-9a-fA-F]{6}").MatchString(value) {
				*output = value
			} else {
				t.input.Reset()
			}
		},
	}
}

func valueOrPlaceholder(input *textinput.Model) string {
	if input.Value() == "" {
		return input.Placeholder
	}
	return input.Value()
}
