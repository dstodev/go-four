package c4_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	assert.Equal(t, expected, actual, "Expected: %v (%#[1]v), Actual: %v (%#[2]v)", expected, actual)
}
