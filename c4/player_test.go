package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
	"github.com/dstodev/go-four/util"
)

func TestPlayerDefault(t *testing.T) {
	var actual c4.Player

	util.AssertEqual(t, c4.None, actual)
}
