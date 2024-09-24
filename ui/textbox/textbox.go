package textbox

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type boxStrategy interface {
	isValid(m Model) bool
	updateOutput(m Model)
}

type Model struct {
	input    textinput.Model
	strategy boxStrategy
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

func NewInteger(output *int, width int) Model {
	input := textinput.New()
	input.Placeholder = strconv.Itoa(*output)
	input.CharLimit = width
	input.Width = width

	return Model{
		input:    input,
		strategy: integerStrategy{output},
	}
}

type integerStrategy struct {
	output *int
}

func (b integerStrategy) isValid(m Model) bool {
	box := &m.input
	value, err := strconv.Atoi(valueOrPlaceholder(box))
	return err == nil && value > 0
}

func (b integerStrategy) updateOutput(m Model) {
	if b.isValid(m) {
		box := &m.input
		*b.output, _ = strconv.Atoi(valueOrPlaceholder(box))
	}
}

func NewString(output *string, width int) Model {
	input := textinput.New()
	input.Placeholder = *output
	input.CharLimit = width
	input.Width = width

	return Model{
		input:    input,
		strategy: stringStrategy{output},
	}
}

type stringStrategy struct {
	output *string
}

func (b stringStrategy) isValid(m Model) bool {
	return true
}

func (b stringStrategy) updateOutput(m Model) {
	if b.isValid(m) {
		box := &m.input
		*b.output = valueOrPlaceholder(box)
	}
}

func NewColor(output *string) Model {
	input := textinput.New()
	input.Placeholder = *output
	input.CharLimit = 6
	input.Width = 6

	return Model{
		input:    input,
		strategy: colorBox{output},
	}
}

type colorBox struct {
	output *string
}

func (b colorBox) isValid(m Model) bool {
	box := &m.input
	value := valueOrPlaceholder(box)
	pattern := regexp.MustCompile("[0-9a-fA-F]{6}")
	return pattern.MatchString(value)
}

func (b colorBox) updateOutput(m Model) {
	if b.isValid(m) {
		box := &m.input
		*b.output = valueOrPlaceholder(box)
	}
}

func (m Model) WithLabel(label string) Model {
	m.SetLabel(label)
	return m
}

func (m *Model) SetLabel(label string) {
	m.input.Prompt = label
}

func (m *Model) Enter() {
	box := &m.input

	box.Reset()
	box.Focus()
	box.CursorEnd()
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
	return m.strategy.isValid(*m)
}

func (m *Model) updateOutput() {
	m.strategy.updateOutput(*m)
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
