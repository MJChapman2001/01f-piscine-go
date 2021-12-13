package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		runnable := true
		res := ""

		for i := range args {
			_, err := strconv.Atoi(args[i])
			if err != nil {
				runnable = false
			} else {
				res += args[i]
			}
		}

		if runnable {
			if args[1] == "/" && args[2] == "0" {
				Zero(args[1])
			} else if args[1] == "%" && args[2] == "0" {
				Zero(args[1])
			} else {
				Run(args)
			}
		}
	}
}

func Run(s []string) int {
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[2])

	if s[1] == "+" {
		return val1 + val2
	} else if s[1] == "-" {
		return val1 - val2
	} else if s[1] == "/" {
		return val1 / val2
	} else if s[1] == "*" {
		return val1 * val2
	} else if s[1] == "%" {
		return val1 % val2
	}

	return 0
}

func Zero(s string) string {
	sign := ""

	if s == "/" {
		sign = "division"
	} else {
		sign = "modulo"
	}

	output := "No " + sign + " by 0"

	return output
}
