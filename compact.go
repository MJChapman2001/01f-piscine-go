package piscine

func Compact(ptr *[]string) int {
	count := 0
	temp := *ptr

	for i := range temp {
		if temp[i] == "" {
			RemoveIndex(temp, i)
		} else {
			count++
		}
	}

	*ptr = temp

	return count
}

func RemoveIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
