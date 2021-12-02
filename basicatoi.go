package BasicAtoi

func BasicAtoi(s string) int {
	val := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "0" {
			val += 0
		} else {
			numArr := []rune(s[i:])
			for j, z := 0, len(numArr); j < len(numArr); j, z = j+1, z-1 {
				pow := 10
				for p := 0; p <= z; p++ {
					if z == 0 {
						pow = 1
					} else {
						pow *= 10
					}
				}
				val += int(numArr[j]) * pow
			}
		}
	}
	return val
}
