package textbox

import (
	"regexp"
	"strconv"
)

type Constraint func(Model) bool

func ConstrainNumeric(m Model) bool {
	box := &m.input
	_, err := strconv.Atoi(valueOrPlaceholder(box))
	return err == nil
}

func ConstrainGreaterZero(m Model) bool {
	return ConstrainGreater(0)(m)
}

func ConstrainGreaterEqZero(m Model) bool {
	return ConstrainGreaterEq(0)(m)
}

func ConstrainGreater(target int) Constraint {
	return func(m Model) bool {
		box := &m.input
		value, err := strconv.Atoi(valueOrPlaceholder(box))
		return err == nil && value > target
	}
}

func ConstrainGreaterEq(target int) Constraint {
	return ConstrainGreater(target - 1)
}

func ConstrainLess(target int) Constraint {
	return func(m Model) bool {
		box := &m.input
		value, err := strconv.Atoi(valueOrPlaceholder(box))
		return err == nil && value < target
	}
}

func ConstrainLessEq(target int) Constraint {
	return ConstrainLess(target + 1)
}

func ConstrainHexColor(m Model) bool {
	box := &m.input
	value := valueOrPlaceholder(box)
	pattern := regexp.MustCompile("[0-9a-fA-F]{6}")
	return pattern.MatchString(value)
}
