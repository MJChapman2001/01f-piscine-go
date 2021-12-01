package PrintComb2

import (
	"github.com/01-edu/z01"
)

func Count(num rune) rune {
	if num == '9' {
		return '0'
	} else {
		return num + 1
	}
}

func PrintComb2() {
	num1 := '0'
	num2 := '0'
	num3 := '0'
	num4 := '0'

	for num1 <= '9' {
		for num2 <= '9' {
			for num3 <= '9' {
				for num4 <= '9' {
					if num1 == num3 && num2 == num4 {
						Count(num4)
					} else {
						z01.PrintRune(num1)
						z01.PrintRune(num2)
						z01.PrintRune(' ')
						z01.PrintRune(num3)
						z01.PrintRune(num4)
						if num1 == '9' && num2 == '8' && num3 == '9' && num4 == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						Count(num4)
					}
				}
				Count(num3)
			}
			Count(num2)
		}
		Count(num1)
	}
}
