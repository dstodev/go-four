package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestPlayerDefault(t *testing.T) {
	var actual c4.Player

	assertEqual(t, c4.None, actual)
}
