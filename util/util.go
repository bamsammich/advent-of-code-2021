package util

import (
	"fmt"
	"sort"
	"strings"
)

func PrintResults(day int, puzzle1, puzzle2 interface{}) {
	fmt.Println(strings.Repeat("-", 44))
	fmt.Printf("| Day %d:%s|\n", day, strings.Repeat(" ", 35))
	fmt.Println(strings.Repeat("-", 44))
	fmt.Printf("| Puzzle 1: %-30v |\n", puzzle1)
	fmt.Printf("| Puzzle 2: %-30v |\n", puzzle2)
	fmt.Println(strings.Repeat("-", 44))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func StringsContain(sl []string, s string) bool {
	for _, str := range sl {
		if str == s {
			return true
		}
	}
	return false
}
