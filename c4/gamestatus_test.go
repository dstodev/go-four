package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestGameStatusDefault(t *testing.T) {
	var actual c4.GameStatus

	util.AssertEqual(t, c4.Initial, actual)
}

func TestGameStatusString(t *testing.T) {
	util.AssertEqual(t, "Initial", c4.Initial.String())
	util.AssertEqual(t, "Running", c4.Running.String())
	util.AssertEqual(t, "Completed", c4.Completed.String())
	util.AssertEqual(t, "Tied", c4.Draw.String())
}
