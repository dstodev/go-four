package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestMin(t *testing.T) {
	util.AssertEqual(t, 1, util.Min(1, 2))
	util.AssertEqual(t, 1, util.Min(2, 1))
	util.AssertEqual(t, 2, util.Min(2, 2))
}
