package piscine

func Compact(ptr *[]string) int {
	temp := *ptr
	var comp []string

	for i, j := 0, 0; i < len(temp); i++ {
		if temp[i] == "" {
			if i != j {
				comp = append(comp, temp[j:i]...)
			}
			j = i + 1
		} else if i == len(temp)-1 {
			comp = append(comp, temp[j:]...)
		}
	}

	*ptr = comp

	return len(comp)
}
