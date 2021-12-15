package piscine

func Compact(ptr *[]string) int {
	var comp []string
	count := 0
	temp := *ptr

	for i := range temp {
		if temp[i] == "" {
			comp = temp[i+1:]
			*ptr = comp
			Compact(ptr)
		} else {
			count++
		}
	}

	return count
}
