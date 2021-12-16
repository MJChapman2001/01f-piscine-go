package piscine

import "sort"

type Pair struct {
	Key   string
	Value int
}

type Pairlist []Pair

func (p Pairlist) Len() int           { return len(p) }
func (p Pairlist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairlist) Less(i, j int) bool { return p[j].Value < p[i].Value }

func MaxWordCountN(text string, n int) map[string]int {
	var textSlice []string

	m := 0
	l := 0
	for l < len(text) {
		if rune(text[l]) == ' ' {
			textSlice = append(textSlice, text[m:l])
			m = l + 1
		} else if l == len(text)-1 {
			textSlice = append(textSlice, text[m:])
		}

		l++
	}

	tempMap := make(map[string]int)

	for i := 0; i < len(textSlice); i++ {
		tempMap[textSlice[i]]++
	}

	sorted := make(Pairlist, len(tempMap))
	index := 0

	for k, v := range tempMap {
		sorted[index] = Pair{k, v}
		index++
	}

	sort.Sort(sorted)

	output := make(map[string]int)

	for i := 0; i < n; i++ {
		j := i + 1
		if j < len(sorted) && sorted[i].Value == sorted[j].Value {
			if sorted[i].Key < sorted[j].Key {
				if i < n-1 {
					output[sorted[i].Key] = sorted[i].Value
					output[sorted[j].Key] = sorted[j].Value
					i++
				} else {
					output[sorted[i].Key] = sorted[i].Value
				}
			} else {
				if len(output) < 2 {
					output[sorted[j].Key] = sorted[j].Value
					output[sorted[i].Key] = sorted[i].Value
					i++
				} else {
					output[sorted[j].Key] = sorted[j].Value
				}
			}
		} else {
			output[sorted[i].Key] = sorted[i].Value
		}
	}

	return output
}
