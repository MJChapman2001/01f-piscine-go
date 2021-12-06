package piscine

import (
	"github.com/01-edu/z01"
)

func PrintReverseAlphabet() {
	alphabet := make([]rune, 26)
	i := 0
	for ch := 'a'; ch <= 'z'; ch++ {
		alphabet[i] = ch
		i++
	}
	for char := 25; char >= 0; char-- {
		z01.PrintRune(alphabet[char])
	}
	z01.PrintRune('\n')
}
