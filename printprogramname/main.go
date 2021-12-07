package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	for i := 2; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}

	z01.PrintRune('\n')
}
