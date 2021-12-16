package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		input, err := strconv.Atoi(args[0])
		var output string

		if err == nil {
			power := 2

			if input == 1 {
				output = "false"
			} else {
				for power < input {
					power *= 2
				}

				if power == input {
					output = "true"
				} else {
					output = "false"
				}
			}
		}

		PrintString(output)
	}
}

func PrintString(s string) {
	for i := range s {
		z01.PrintRune(rune(s[i]))
	}

	z01.PrintRune('\n')
}
