package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	if len(name) == 0 {
		z01.PrintRune(rune(0))
	} else {
		for i := 0; i < len(name); i++ {
			if IsAlpha(name[i]) {
				z01.PrintRune(rune(name[i]))
			}
		}
	}

	z01.PrintRune('\n')
}

func IsAlpha(s byte) bool {
	if s < 48 || s > 122 {
		return false
	} else if s > 57 && s < 65 {
		return false
	} else if s > 90 && s < 97 {
		return false
	}
	return true
}
