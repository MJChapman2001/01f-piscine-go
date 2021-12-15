package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		runnable := true

		_, err1 := strconv.Atoi(args[0])
		_, err2 := strconv.Atoi(args[2])

		if err1 != nil || err2 != nil {
			runnable = false
		}

		if runnable {
			os.Stdout.WriteString(Maths(args))
		}
	}
}

func Maths(s []string) string {
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[2])
	var result int
	var output string

	switch s[1] {
	case "+":
		result = val1 + val2
		output = IntToString(result)
	case "-":
		result = val1 - val2
		output = IntToString(result)
	case "*":
		result = val1 * val2
		output = IntToString(result)
	case "/":
		if val2 == 0 {
			return "No division by 0"
		} else {
			result = val1 / val2
			output = IntToString(result)
		}
	case "%":
		if val2 == 0 {
			return "No modulo by 0"
		} else {
			result = val1 % val2
			output = IntToString(result)
		}
	}

	return output
}

func IntToString(nbr int) string {
	var runeArr []rune
	output := ""

	for nbr != 0 {
		runeArr = append(runeArr, rune((nbr%10)+48))
		nbr /= 10
	}

	for i := len(runeArr) - 1; i >= 0; i-- {
		output += string(runeArr[i])
	}

	return output
}
