package util_test

import (
	"testing"

	"github.com/dstodev/go-four/util"
)

func TestLastN(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	last3 := util.LastN(elements, 3)
	util.AssertEqual(t, []int{3, 4, 5}, last3)
}

func TestLastNOverrun(t *testing.T) {
	elements := []int{1}
	last3 := util.LastN(elements, 3)
	util.AssertEqual(t, []int{1}, last3)
}

func TestLastNEmpty(t *testing.T) {
	elements := []int{}
	last3 := util.LastN(elements, 3)
	util.AssertEqual(t, []int{}, last3)
}
