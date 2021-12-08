package piscine

func AppendRange(min, max int) []int {
	var result []int

	if min >= max {
		return result
	}

	for min < max {
		result = append(result, min)
		min++
	}

	return result
}
