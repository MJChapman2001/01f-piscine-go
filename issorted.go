package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	val := 0
	pos := 0

	for val == 0 {
		for pos < len(a)-1 {
			val = f(a[pos], a[pos+1])
			pos++
		}
	}

	if val > 0 {
		for i := pos; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	} else if val < 0 {
		for i := pos; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}

	return true
}
