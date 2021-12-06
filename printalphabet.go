package piscine

import (
	"github.com/01-edu/z01"
)

func PrintAlphabet() {
	for ch := 'a'; ch <= 'z'; ch++ {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
