package util

func LastN(elements []int, n int) []int {
	if n > len(elements) {
		return elements
	}
	return elements[len(elements)-n:]
}
