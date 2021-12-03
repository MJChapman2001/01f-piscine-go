package BasicAtoi

func BasicAtoi(s string) int {
	runeArr := []rune(s)
	val := 0
	for i, j := 0, len(runeArr); i < len(runeArr); i, j = i+1, j-1 {
		pow := 10
		if int(runeArr[i])-48 == 0 {
			val += 0
		} else {
			for z := 1; z < j; z++ {
				pow *= 10
			}
			val += (int(runeArr[i]) - 48) * pow
		}
	}
	return val
}
