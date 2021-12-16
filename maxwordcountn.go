package piscine

import "sort"

type Pair struct {
	Key   string
	Value int
}

type Pairlist []Pair

func (p Pairlist) Len() int           { return len(p) }
func (p Pairlist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairlist) Less(i, j int) bool { return p[i].Value < p[j].Value }

func MaxWordCountN(text string, n int) map[string]int {
	var textSlice []string

	for i, j := 0, 0; i < len(text); i++ {
		if rune(text[i]) == ' ' || rune(text[i]) == '\n' {
			if i != j {
				textSlice = append(textSlice, text[j:i])
			}
			j = i + 1
		} else if i == len(text)-1 {
			textSlice = append(textSlice, text[j:])
		}
	}

	tempMap := make(map[string]int)

	for _, i := range textSlice {
		if tempMap[i] > 0 {
			tempMap[i]++
		} else {
			tempMap[i] = 1
		}
	}

	sorted := make(Pairlist, len(tempMap))
	index := 0

	for k, v := range tempMap {
		sorted[index] = Pair{k, v}
		index++
	}

	sort.Sort(sorted)

	output := make(map[string]int, n)

	for i := 0; i < n; i++ {
		output[sorted[i].Key] = sorted[i].Value
	}

	return output
}
