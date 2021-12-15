package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		runnable := true

		for i := range args {
			_, err := strconv.Atoi(args[i])
			if err != nil {
				runnable = false
			}
		}

		if runnable {
			Maths(args)
		}
	}
}

func Maths(s []string) int {
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[2])

	switch s[1] {
	case "+":
		return val1 + val2
	case "-":
		return val1 - val2
	case "*":
		return val1 * val2
	case "/":
		if val2 == 0 {
			Zero(s[1])
		} else {
			return val1 / val2
		}
	case "%":
		if val2 == 0 {
			Zero(s[1])
		} else {
			return val1 % val2
		}
	}

	return 0
}

func Zero(sign string) string {
	if sign == "/" {
		return "No division by 0"
	} else {
		return "No modulo by 0"
	}
}
