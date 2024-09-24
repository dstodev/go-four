package textbox

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	input       textinput.Model
	update      func(m Model)
	constraints []Constraint
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.input.View()
}

func NewInteger(output *int, width int, constraints ...Constraint) Model {
	input := textinput.New()
	input.Placeholder = strconv.Itoa(*output)
	input.CharLimit = width
	input.Width = width

	return Model{
		input: input,
		update: func(m Model) {
			if m.isValid() {
				box := &m.input
				*output, _ = strconv.Atoi(valueOrPlaceholder(box))
			}
		},
		constraints: constraints,
	}
}

func NewString(output *string, width int, constraints ...Constraint) Model {
	input := textinput.New()
	input.Placeholder = *output
	input.CharLimit = width
	input.Width = width

	return Model{
		input: input,
		update: func(m Model) {
			if m.isValid() {
				box := &m.input
				*output = valueOrPlaceholder(box)
			}
		},
		constraints: constraints,
	}
}

func NewColor(output *string, constraints ...Constraint) Model {
	input := textinput.New()
	input.Placeholder = *output
	input.CharLimit = 6
	input.Width = 6

	constraints = append(constraints, ConstrainHexColor)

	return Model{
		input: input,
		update: func(m Model) {
			if m.isValid() {
				box := &m.input
				*output = valueOrPlaceholder(box)
			}
		},
		constraints: constraints,
	}
}

func (m Model) WithLabel(label string) Model {
	m.SetLabel(label)
	return m
}

func (m *Model) SetLabel(label string) {
	m.input.Prompt = label
}

func (m *Model) Enter() tea.Cmd {
	box := &m.input

	box.Reset()
	cmd := box.Focus()
	box.CursorEnd()

	return cmd
}

func (m *Model) Leave(assertDifferent *Model) {
	box := &m.input
	box.Blur()

	if !m.isValid() {
		box.Reset()
	}

	if assertDifferent != nil {
		otherBox := &assertDifferent.input

		equalToOther := func() bool {
			return strings.EqualFold(valueOrPlaceholder(box), valueOrPlaceholder(otherBox))
		}

		if equalToOther() {
			box.Reset()

			// If resetting causes a conflict (because the other box is using this box's placeholder)
			if equalToOther() {
				box.SetValue(otherBox.Placeholder)
			}
		}
	}

	m.updateOutput()
}

func (m *Model) isValid() bool {
	for _, constraint := range m.constraints {
		if !constraint(*m) {
			return false
		}
	}
	return true
}

func (m *Model) updateOutput() {
	m.update(*m)
}

func (m *Model) Value() string {
	return valueOrPlaceholder(&m.input)
}

func valueOrPlaceholder(box *textinput.Model) string {
	value := box.Value()

	if value == "" {
		return box.Placeholder
	}
	return value
}
