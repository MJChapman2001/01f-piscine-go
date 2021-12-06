package piscine

func IterativeFactorial(nb int) int {
	for i := 1; i < nb; i++ {
		nb *= i
	}

	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		return nb
	}
}
