package piscine

func IterativeFactorial(nb int) int {
	result := 0

	if nb > 0 {
		result = 1
		for i := nb; i > 1; i-- {
			result *= i
		}
	}

	return result
}
