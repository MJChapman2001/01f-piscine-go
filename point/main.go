package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(prt *point) {
	prt.x = 42
	prt.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	output := ("x = !, y = *")

	for i := range output {
		if output[i] == '!' {
			j := points.x % 10
			k := (points.x - j) / 10

			z01.PrintRune(rune(k + 48))
			z01.PrintRune(rune(j + 48))
		} else if output[i] == '*' {
			j := points.y % 10
			k := (points.y - j) / 10

			z01.PrintRune(rune(k + 48))
			z01.PrintRune(rune(j + 48))
		} else {
			z01.PrintRune(rune(output[i]))
		}
	}

	z01.PrintRune('\n')
}
