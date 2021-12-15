package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var message string

	if len(args) > 0 {
		for i := range args {
			file, err := ioutil.ReadFile(args[i])

			if err != nil {
				message = "ERROR: " + err.Error()
			} else {
				message = string(file)
			}

			for j := range message {
				z01.PrintRune(rune(message[j]))
			}

			if err != nil {
				z01.PrintRune('\n')
				os.Exit(1)
			}
		}
	} else if len(args) == 0 {
		message, err := ioutil.ReadAll(os.Stdin)

		if err == nil {
			for j := range message {
				z01.PrintRune(rune(message[j]))
			}
		}
	}
}
