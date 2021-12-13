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
			z01.PrintRune(rune(points.x))
		} else if output[i] == '*' {
			z01.PrintRune(rune(points.y))
		} else {
			z01.PrintRune(rune(output[i]))
		}
	}

	z01.PrintRune('\n')
}
