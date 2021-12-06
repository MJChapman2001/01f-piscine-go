package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		found := false
		if len(toFind) == 0 {
			break
		}
		if s[i] == toFind[0] {
			for j := 0; j < len(toFind); j++ {
				if s[i+j] == toFind[j] {
					found = true
				} else {
					found = false
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
