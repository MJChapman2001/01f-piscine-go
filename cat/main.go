package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for i := range args {
			file, err := ioutil.ReadFile(args[i])
			var message string

			if err != nil {
				message = err.Error()
			} else {
				message = string(file)
			}

			for j := range message {
				z01.PrintRune(rune(message[j]))
			}

			if message == err.Error() {
				os.Exit(1)
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}
