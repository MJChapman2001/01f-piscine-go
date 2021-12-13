package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	val := 0

	for val == 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) != 0 {
				val = f(a[i], a[i+1])
				break
			}
		}
	}

	if val > 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	} else if val < 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}

	return true
}
