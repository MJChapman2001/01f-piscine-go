package StrLen

func StrLen(s string) int {
	runeArr := make([]rune, 0)

	for i := 0; i < len(s); i++ {
		runeArr = append(runeArr, rune(s[i]))
	}
	return len(runeArr)
}
