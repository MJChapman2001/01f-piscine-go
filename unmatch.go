package piscine

func Unmatch(a []int) int {
	val := -1

	for i := range a {
		found := false

		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				found = true
				break
			}
		}

		if !found {
			val = a[i]
			break
		}
	}

	return val
}
