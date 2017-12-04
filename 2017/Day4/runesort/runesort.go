package runesort

import (
	"sort"
)

type runes []rune

// Sort reorders a string alphabetically.
func Sort(s string) string {
	r := []rune(s)
	sort.Sort(runes(r))
	return string(r)
}

func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r runes) Len() int {
	return len(r)
}
