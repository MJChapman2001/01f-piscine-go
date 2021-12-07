package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	var numArr []int
	var runeArr []rune

	if len(params) > 0 {
		if params[0] == "--upper" {
			for i := 1; i < len(params); i++ {
				numArr = append(numArr, BasicAtoi2(params[i]))
			}

			for j := 0; j < len(numArr); j++ {
				if numArr[j] == 0 || numArr[j] > 26 {
					runeArr = append(runeArr, ' ')
				} else {
					runeArr = append(runeArr, ('A' + rune(numArr[j]-1)))
				}
			}
		} else {
			for i := 0; i < len(params); i++ {
				numArr = append(numArr, BasicAtoi2(params[i]))
			}

			for j := 0; j < len(numArr); j++ {
				if numArr[j] == 0 || numArr[j] > 26 {
					runeArr = append(runeArr, ' ')
				} else {
					runeArr = append(runeArr, ('a' + rune(numArr[j]-1)))
				}
			}
		}

		for k := 0; k < len(runeArr); k++ {
			z01.PrintRune(runeArr[k])
		}
	}

	z01.PrintRune('\n')
}

func BasicAtoi2(s string) int {
	runeArr := []rune(s)
	val := 0
	for i, j := 0, len(runeArr); i < len(runeArr); i, j = i+1, j-1 {
		pow := 1
		if int(runeArr[i]) < 48 || int(runeArr[i]) > 57 {
			val = 0
			break
		} else {
			for z := 1; z < j; z++ {
				pow *= 10
			}
			val += (int(runeArr[i]) - 48) * pow
		}
	}
	return val
}
