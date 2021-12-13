package piscine

func Any(f func(string) bool, a []string) bool {
	for i := range a {
		if f(a[i]) {
			return true
		}
	}

	return false
}
