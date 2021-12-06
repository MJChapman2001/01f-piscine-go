package piscine

func Index(s string, toFind string) int {
	if len(toFind) > 1 {
		for i := 0; i < len(s); i++ {
			if s[i:(i+(len(toFind)))] == toFind {
				return i
			}
		}
	}

	if len(toFind) == 1 {
		for i := 0; i < len(s); i++ {
			if string(s[i]) == toFind {
				return i
			}
		}
	}
	return -1
}
