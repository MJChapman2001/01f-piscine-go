package piscine

func Itoa(n int) string {
	output := ""
	var arr []rune

	if n < 0 {
		output += "-"
		n *= -1
	}

	for n != 0 {
		temp := n % 10
		arr = append(arr, rune(temp+48))
		n -= temp
		n /= 10
	}

	for i := len(arr) - 1; i >= 0; i-- {
		output += string(arr[i])
	}

	return output
}
