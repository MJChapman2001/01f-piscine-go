package StrLen

func StrLen(s string) int {
	count := 0
	for i := 0; i <= len(s); i++ {
		count += 1
	}
	return count
}
