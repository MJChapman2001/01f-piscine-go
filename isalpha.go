package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 122 {
			return false
		} else if s[i] > 57 && s[i] < 65 {
			return false
		} else if s[i] > 90 && s[i] < 97 {
			return false
		}
	}
	return true
}
