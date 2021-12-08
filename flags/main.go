package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var output []rune

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Println("--insert\n   -i\n   This flag inserts the string into the string passed as argument.")
		fmt.Println("--order\n   -o\n   This flag will behave like a boolean, if it is called it will order the argument.")
	} else {
		for i := 0; i < len(args); i++ {
			if len(args[i]) > 9 && args[i][0:9] == "--insert=" {
				for j := 9; j < len(args[i]); j++ {
					output = append(output, rune(args[i][j]))
				}
			} else if len(args[i]) > 3 && args[i][0:3] == "-i=" {
				for j := 3; j < len(args[i]); j++ {
					output = append(output, rune(args[i][j]))
				}
			} else if args[i] == "--order" || args[i] == "-o" && i < len(args)-1 {
				for j := 0; j < len(args[i+1]); j++ {
					output = append(output, rune(args[i+1][j]))
					Sort(output)
				}
				i++
				break
			} else {
				runeArr := []rune(args[i])
				output = append(runeArr, output...)
			}
		}

		for k := 0; k < len(output); k++ {
			z01.PrintRune(output[k])
		}

		z01.PrintRune('\n')
	}
}

func Sort(n []rune) []rune {
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
