package util

import (
	"strings"
)

func LastNLines(s string, n int) string {
	tokens := strings.Split(s, "\n")

	if n >= len(tokens) {
		return s
	}
	return strings.Join(tokens[len(tokens)-n:], "\n")
}
