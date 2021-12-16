package piscine

func Unmatch(a []int) int {
	val := -1

	for i := range a {
		count := 1

		for j := 0; j < len(a); j++ {
			if a[i] == a[j] && i != j {
				count++
			}
		}

		if count%2 != 0 {
			val = a[i]
			break
		}
	}

	return val
}
