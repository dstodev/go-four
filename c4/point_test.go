package c4_test

import (
	"testing"

	"github.com/dstodev/go-four/c4"
)

func TestPointDestructure(t *testing.T) {
	point := c4.Point{1, 2}

	r, c := point.Get()

	assertEqual(t, 1, r)
	assertEqual(t, 2, c)
}
