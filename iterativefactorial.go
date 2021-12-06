package piscine

func IterativeFactorial(nb int) int {
	result := 1

	for i := nb; i > 1; i-- {
		result *= i
	}

	return result
}
