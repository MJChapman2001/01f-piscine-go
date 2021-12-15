package piscine

func ActiveBits(n int) int {
	count := 0

	for i := 128; i >= 1; i /= 2 {
		if n > i {
			n -= i
			count++
		}
	}

	return count
}
