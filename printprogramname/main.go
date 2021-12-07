package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1]

	for i := 0; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}
}
