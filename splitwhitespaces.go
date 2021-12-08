package piscine

func SplitWhiteSpace(s string) []string {
	var output []string

	for i, j := 0, 0; i < len(s); i++ {
		if rune(s[i]) == ' ' || rune(s[i]) == '\n' {
			output = append(output, s[j:i])
			j = i + 1
		}
	}

	return output
}
