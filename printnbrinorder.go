package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nArr []int

	for n != 0 {
		nArr = append(nArr, n%10)
		n /= 10
	}

	Sort(nArr)

	for k := 0; k < len(nArr); k++ {
		z01.PrintRune(rune(nArr[k]))
	}
	z01.PrintRune('\n')
}

func Sort(n []int) []int {
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
