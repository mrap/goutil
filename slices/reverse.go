package slices

func ReverseInts(a []int) {
	len := len(a)
	for i := 0; i < len/2; i++ {
		j := len - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}

func ReverseStrings(a []string) {
	len := len(a)
	for i := 0; i < len/2; i++ {
		j := len - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}
