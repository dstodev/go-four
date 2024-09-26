package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestLastNLines(t *testing.T) {
	elements := "a\nb\nc\nd\ne"
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, "c\nd\ne", last3)
}

func TestLastNLinesOverrun(t *testing.T) {
	elements := "a"
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, "a", last3)
}

func TestLastNLinesEmpty(t *testing.T) {
	elements := ""
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, "", last3)
}

func TestLastNLinesPreservesNewlines(t *testing.T) {
	elements := "a\nb\nc\nd\n\ne"
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, "d\n\ne", last3)
}
