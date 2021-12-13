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
			for j := points.x; j != 0; j /= 10 {
				z01.PrintRune(rune((j % 10) + 48))
			}
		} else if output[i] == '*' {
			for j := points.y; j != 0; j /= 10 {
				z01.PrintRune(rune((j % 10) + 48))
			}
		} else {
			z01.PrintRune(rune(output[i]))
		}
	}

	z01.PrintRune('\n')
}
