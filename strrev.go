package StrRev

func StrRev(s string) string {
	runeArr := []rune(s)
	reverse := ""
	for i := len(runeArr) - 1; i > 0; i++ {
		reverse += string(runeArr[i])
	}
	return reverse
}
