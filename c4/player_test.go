package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestPlayerDefault(t *testing.T) {
	var actual c4.Player

	assertEqual(t, c4.None, actual)
}

func TestPlayerString(t *testing.T) {
	assertEqual(t, "None", c4.None.String())
	assertEqual(t, "One", c4.One.String())
	assertEqual(t, "Two", c4.Two.String())
}

func TestPlayerNegate(t *testing.T) {
	assertEqual(t, c4.None, c4.None.Negate()) // None.Negate() == None
	assertEqual(t, c4.Two, c4.One.Negate())
	assertEqual(t, c4.One, c4.Two.Negate())
}

func TestPlayerShort(t *testing.T) {
	assertEqual(t, " ", c4.None.Short().String())
	assertEqual(t, "A", c4.One.Short().String())
	assertEqual(t, "B", c4.Two.Short().String())
}
