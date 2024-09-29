package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestLastNLines0(t *testing.T) {
	elements := "a\nb\nc"
	last0 := util.LastNLines(elements, 0)
	util.AssertEqual(t, "", last0)
}

func TestLastNLines1(t *testing.T) {
	elements := "a\nb\nc"
	last1 := util.LastNLines(elements, 1)
	util.AssertEqual(t, "c", last1)
}

func TestLastNLines2(t *testing.T) {
	elements := "a\nb\nc"
	last2 := util.LastNLines(elements, 2)
	util.AssertEqual(t, "b\nc", last2)
}

func TestLastNLines3(t *testing.T) {
	elements := "a\nb\nc"
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, elements, last3)
}

func TestLastNLines4Overrun(t *testing.T) {
	elements := "a\nb\nc"
	last4 := util.LastNLines(elements, 4)
	util.AssertEqual(t, elements, last4)
}

func TestLastNLinesEmpty(t *testing.T) {
	elements := ""
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, "", last3)
}
func TestLastNLinesPreservesNewlines(t *testing.T) {
	elements := "a\nb\nc\nd\n\ne\n"
	last5 := util.LastNLines(elements, 5)
	util.AssertEqual(t, "c\nd\n\ne\n", last5)
	// c\n
	// d\n
	// \n
	// e\n
	//
}
