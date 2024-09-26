package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestMax(t *testing.T) {
	util.AssertEqual(t, 2, util.Max(1, 2))
	util.AssertEqual(t, 2, util.Max(2, 1))
	util.AssertEqual(t, 1, util.Max(1, 1))
}
