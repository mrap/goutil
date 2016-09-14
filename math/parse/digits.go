package parse

import "github.com/mrap/goutil/slices"

func IntDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n < 0 {
		n *= -1
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	slices.ReverseInts(digits)
	return digits
}
