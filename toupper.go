package piscine

func ToUpper(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			result += string(s[i] - 32)
		} else {
			result += string(s[i])
		}
	}
	return result
}
