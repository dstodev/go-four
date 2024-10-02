package util

import (
	"strings"
)

func LastNLines(s string, n int) []string {
	tokens := strings.Split(s, "\n")

	if len(tokens) == 1 {
		if tokens[0] == "" {
			return []string{}
		}
		return tokens
	}
	if n >= len(tokens) {
		return tokens
	}
	return tokens[len(tokens)-n:] // n==0 is ok
}
