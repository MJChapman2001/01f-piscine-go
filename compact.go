package piscine

func Compact(ptr *[]string) int {
	count := 0
	temp := *ptr

	for i := range temp {
		if temp[i] == "" {
			*ptr = temp[i+1:]
			Compact(ptr)
		} else {
			count++
		}
	}

	return count
}
