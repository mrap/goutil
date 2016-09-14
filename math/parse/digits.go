package parse

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

	len := len(digits)
	for i := 0; i < len/2; i++ {
		j := len - 1 - i
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}
