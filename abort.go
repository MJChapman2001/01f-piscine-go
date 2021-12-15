package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	SortInt(arr)

	return arr[2]
}

func SortInt(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			temp := a[i]
			a[i] = a[i+1]
			a[i+1] = temp
			Sort(a)
		}
	}

	return a
}
