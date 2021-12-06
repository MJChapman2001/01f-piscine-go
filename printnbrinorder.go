package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nArr []int

	for n != 0 {
		nArr = append(nArr, n%10)
		n /= 10
	}

	for j := 0; j < len(nArr); j++ {
		Sort(nArr)
	}

	for k := 0; k < len(nArr); k++ {
		z01.PrintRune(rune(nArr[k]))
	}
}

func Sort(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			a := n[i]
			n[i] = n[i+1]
			n[i+1] = a
		}
	}

	return n
}
