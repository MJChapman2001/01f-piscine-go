package piscine

func Compact(ptr *[]string) int {
	temp := *ptr
	var comp []string

	for i := range temp {
		if temp[i] == "" {
			comp = RemoveIndex(temp, i)
		}
	}

	*ptr = comp

	return len(comp)
}

func RemoveIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
