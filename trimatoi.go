package piscine

func TrimAtoi(s string) int {
	result := 0
	sArr := []rune(s)

	for i := 0; i < len(sArr); i++ {
		if IsNumeric(string(s[i])) {
			result *= 10
			result += int(sArr[i]) - 48
		}
	}

	return result
}
