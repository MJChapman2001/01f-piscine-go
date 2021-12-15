package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			temp := a[i]
			a[i] = a[i+1]
			a[i+1] = temp
			SortWordArr(a)
		}
	}
}
