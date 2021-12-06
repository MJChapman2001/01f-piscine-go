package piscine

func LastRune(s string) rune {
	runeArr := []rune(s)
	return runeArr[len(runeArr)-1]
}
