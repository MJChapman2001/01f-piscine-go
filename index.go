package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		found := false
		if s[i] == toFind[0] {
			for j := 0; j < len(toFind); j++ {
				if i+j < len(s) {
					if s[i+j] == toFind[j] {
						found = true
					} else {
						found = false
					}
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
