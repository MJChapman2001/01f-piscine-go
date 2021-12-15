package piscine

func Compact(ptr *[]string) int {
	temp := *ptr

	for i := range temp {
		if temp[i] == "" {
			temp = RemoveIndex(temp, i)
		}
	}

	*ptr = temp

	return len(temp)
}

func RemoveIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
