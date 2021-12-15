package piscine

func Compact(ptr *[]string) int {
	count := 0
	temp := *ptr
	comp := make([]string, len(temp))

	for i := range temp {
		if temp[i] != "" {
			count++
			for j := range comp {
				if comp[j] == "" {
					comp[j] = temp[i]
				}
			}
		}
	}

	for k := range comp {
		if comp[k] == "" {
			comp = comp[0:k]
		}
	}

	*ptr = comp

	return count
}
