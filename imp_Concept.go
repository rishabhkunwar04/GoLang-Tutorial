//### TO SORT STRING IN GOLANG

package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	w1 := "bcad"
	w2 := SortString(w1)

	fmt.Println(w1)
	fmt.Println(w2)
}

//

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

func main() {
	w1 := "bcad"
	w2 := SortString(w1)

	fmt.Println(w1)
	fmt.Println(w2)
}
