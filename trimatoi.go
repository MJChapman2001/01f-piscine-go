package piscine

func TrimAtoi(s string) int {
	result := 0
	sArr := []rune(s)
	negative := false

	for i := 0; i < len(sArr); i++ {
		if s[i] == 45 && result == 0 {
			negative = true
		}
		if IsNumeric(string(s[i])) {
			result *= 10
			result += int(sArr[i]) - 48
		}
	}

	if negative {
		return -result
	} else {
		return result
	}
}
