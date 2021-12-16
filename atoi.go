package piscine

func Atoi(s string) int {
	output := 0
	sign := 1

	for i := range s {
		if s[i] == '-' && output == 0 {
			sign = -1
		}

		if s[i] > 47 && s[i] < 58 {
			output *= 10
			output += int(s[i] - 48)
		} else {
			output = 0
			break
		}
	}

	return output * sign
}
