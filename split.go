package piscine

func Split(s, sep string) []string {
	var output []string

	for i, j := 0, 0; i < len(s); i++ {
		if i+len(sep) < len(s)-1 {
			if s[i:i+len(sep)] == sep {
				output = append(output, s[j:i])
				j = i + len(sep)
			}
		} else if i == len(s)-1 {
			output = append(output, s[j:])
		}
	}

	return output
}
