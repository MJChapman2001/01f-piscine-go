package PrintComb

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for num1 := '0'; num1 < '8'; num1++ {
		for num2 := num1 + 1; num2 < '9'; num2++ {
			for num3 := num2 + 1; num3 <= '9'; num3++ {
				z01.PrintRune(num1)
				z01.PrintRune(num2)
				z01.PrintRune(num3)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
