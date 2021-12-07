package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]

	Sort(name)

	for i := 0; i < len(name); i++ {
		runeArr := []rune(name[i])

		for j := 0; j < len(runeArr); j++ {
			z01.PrintRune(runeArr[j])
		}

		z01.PrintRune('\n')
	}
}

func Sort(n []string) []string {
	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			a := n[i]
			n[i] = n[i+1]
			n[i+1] = a
		}
	}

	for j := 0; j < len(n)-1; j++ {
		if n[j] > n[j+1] {
			Sort(n)
		}
	}

	return n
}
