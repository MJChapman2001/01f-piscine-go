package PrintComb

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	firstDigit := '0'
	secondDigit := '0'
	thirdDigit := '0'

	for firstDigit < '8' {
		secondDigit = firstDigit + 1
		thirdDigit = secondDigit + 1

		for secondDigit <= '9' {

			for thirdDigit <= '9' {
				z01.PrintRune(firstDigit)
				z01.PrintRune(secondDigit)
				z01.PrintRune(thirdDigit)
				z01.PrintRune(',')
				z01.PrintRune(' ')
				thirdDigit++
			}
			secondDigit++
		}
		firstDigit++
	}
}
