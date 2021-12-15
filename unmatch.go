package piscine

func Unmatch(a []int) int {
	val := -1

	for i := range a {
		found := false

		for j := range a[i+1:] {
			if a[i] == a[j] {
				found = true
			}
		}

		if !found {
			val = a[i]
			break
		}
	}

	return val
}
