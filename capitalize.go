package piscine

func Capitalize(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if i == 0 || !IsAlpha(string(s[i-1])) {
			if !IsNumeric(string(s[i])) && IsLower(string(s[i])) {
				result += ToUpper(string(s[i]))
			} else {
				result += string(s[i])
			}
		} else if IsAlpha(string(s[i])) {
			if !IsNumeric(string(s[i])) && IsUpper(string(s[i])) {
				result += ToLower(string(s[i]))
			} else {
				result += string(s[i])
			}
		} else {
			result += string(s[i])
		}
	}
	return result
}
