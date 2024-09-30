package util

import "cmp"

func Max[T cmp.Ordered](a, b T) T {
	if a == Min(a, b) {
		return b
	}
	return a
}
