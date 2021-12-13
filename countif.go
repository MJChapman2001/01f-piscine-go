package piscine

func CountIf(f func(string) bool, tab []string) int {
	val := 0

	for i := range tab {
		if f(tab[i]) {
			val++
		}
	}

	return val
}
