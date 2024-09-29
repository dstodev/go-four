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

func TestLastNLines0(t *testing.T) {
	elements := "a\nb\nc\nd\ne"
	last0 := util.LastNLines(elements, 0)
	util.AssertEqual(t, "", last0)
}

func TestLastNLines1(t *testing.T) {
	elements := "a\nb\nc\nd\ne"
	last1 := util.LastNLines(elements, 1)
	util.AssertEqual(t, "e", last1)
}

func TestLastNLines5(t *testing.T) {
	elements := "a\nb\nc\nd\ne"
	last5 := util.LastNLines(elements, 5)
	util.AssertEqual(t, elements, last5)
}

func TestLastNLinesOverrun(t *testing.T) {
	elements := "a\nb"
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, "a\nb", last3)
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
