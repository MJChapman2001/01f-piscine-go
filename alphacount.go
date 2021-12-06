package piscine

func AlphaCount(s string) int {
	runeArr := []rune(s)
	count := 0

	for i := 0; i < len(runeArr); i++ {
		if runeArr[i] >= 65 && runeArr[i] <= 90 {
			count++
		} else if runeArr[i] >= 97 && runeArr[i] <= 122 {
			count++
		}
	}

	return count
}
