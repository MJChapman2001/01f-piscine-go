package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	result := nb
	for i := 1; i < power; i++ {
		result *= nb
	}

	return result
}
