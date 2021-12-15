package piscine

func Rot14(s string) string {
	arr := []rune(s)
	output := ""

	for i := range arr {
		if s[i] > 64 && s[i] < 91 {
			if s[i]+14 > 90 {
				arr[i] = rune((s[i] + 14) - (90 - 64))
			} else {
				arr[i] = rune(s[i] + 14)
			}
		} else if s[i] > 96 && s[i] < 123 {
			if s[i]+14 > 122 {
				arr[i] = rune((s[i] + 14) - (122 - 96))
			} else {
				arr[i] = rune(s[i] + 14)
			}
		}

		output += string(arr[i])
	}

	return output
}
