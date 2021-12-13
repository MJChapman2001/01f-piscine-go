package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for i := range a {
		result = append(result, f(a[i]))
	}

	return result
}
