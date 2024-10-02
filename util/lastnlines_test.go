package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestLastNLines0(t *testing.T) {
	elements := "a\nb\nc"
	last0 := util.LastNLines(elements, 0)
	util.AssertEqual(t, []string{}, last0)
}

func TestLastNLines1(t *testing.T) {
	elements := "a\nb\nc"
	last1 := util.LastNLines(elements, 1)
	util.AssertEqual(t, []string{"c"}, last1)
}

func TestLastNLines2(t *testing.T) {
	elements := "a\nb\nc"
	last2 := util.LastNLines(elements, 2)
	util.AssertEqual(t, []string{"b", "c"}, last2)
}

func TestLastNLines3(t *testing.T) {
	elements := "a\nb\nc"
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, []string{"a", "b", "c"}, last3)
}

func TestLastNLines4Overrun(t *testing.T) {
	elements := "a\nb\nc"
	last4 := util.LastNLines(elements, 4)
	util.AssertEqual(t, []string{"a", "b", "c"}, last4)
}

func TestLastNLinesEmpty(t *testing.T) {
	elements := ""
	last3 := util.LastNLines(elements, 3)
	util.AssertEqual(t, []string{}, last3)
}

func TestLastNLinesOneLineNoNewline(t *testing.T) {
	elements := "a"
	last1 := util.LastNLines(elements, 1)
	util.AssertEqual(t, []string{"a"}, last1)
}

func TestLastNLinesPreservesNewlines(t *testing.T) {
	elements := "\na\n\nb\n"
	last5 := util.LastNLines(elements, 5)
	util.AssertEqual(t, []string{"", "a", "", "b", ""}, last5)
	// 1. \n
	// 2. a\n
	// 3. \n
	// 4. b\n
	// 5.
}
