package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := 0; i < 9900; i++ {
		num4 := i % 10
		num3 := ((i % 100) - num4) / 10
		num2 := ((i % 1000) - (num4 + num3)) / 100
		num1 := (i - (num4 + num3 + num2)) / 1000

		z01.PrintRune(rune(num1))
		z01.PrintRune(rune(num2))
		z01.PrintRune(' ')
		z01.PrintRune(rune(num3))
		z01.PrintRune(rune(num4))

		if num1 == '9' && num2 == '8' && num3 == '9' && num4 == '9' {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
