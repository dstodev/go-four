package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	assert.Equal(t, expected, actual, "Expected: %v (%#[1]v), Actual: %v (%#[2]v)", expected, actual)
}
