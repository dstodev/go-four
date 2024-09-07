package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestGameStatusDefault(t *testing.T) {
	var actual c4.GameStatus

	assertEqual(t, c4.Initial, actual)
}

func TestGameStatusString(t *testing.T) {
	assertEqual(t, "Initial", c4.Initial.String())
	assertEqual(t, "Running", c4.Running.String())
	assertEqual(t, "Completed", c4.Completed.String())
}
