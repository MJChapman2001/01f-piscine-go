package piscine

func NRune(s string, n int) rune {
	runeArr := []rune(s)

	if n <= 0 || n >= len(runeArr) {
		return 0
	} else {
		return runeArr[n-1]
	}
}
