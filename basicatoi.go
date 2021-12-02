package BasicAtoi

func BasicAtoi(s string) int {
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && string(s[i]) == "0" {
			return 0
		} else if string(s[i]) == "0" {
			continue
		} else {
			numArr := []rune(s[i:])
			val := 0
			for j, z := 0, len(numArr); j < len(numArr); j, z = j+1, z-1 {
				pow := 10
				for p := 0; p <= z; p++ {
					if z == 0 {
						pow = 1
					} else {
						pow *= pow
					}
				}
				val += int(numArr[j]) * pow
			}
			return val
		}
	}
}
