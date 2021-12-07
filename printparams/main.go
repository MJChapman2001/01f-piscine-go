package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]

	for i := 0; i < len(name); i++ {
		runeArr := []rune(name[i])

		for j := 0; j < len(runeArr); j++ {
			z01.PrintRune(runeArr[j])
		}

		z01.PrintRune('\n')
	}
}
