package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestCountLines0(t *testing.T) {
	actual := util.CountLines("")
	util.AssertEqual(t, 0, actual)
}

func TestCountLines1(t *testing.T) {
	actual := util.CountLines("Hello, World!")
	util.AssertEqual(t, 1, actual)
}

func TestCountLines2(t *testing.T) {
	actual := util.CountLines("Hello,\nWorld!")
	util.AssertEqual(t, 2, actual)
}

func TestCountLines3(t *testing.T) {
	actual := util.CountLines("Hello,\n\nWorld!")
	util.AssertEqual(t, 3, actual)
}
