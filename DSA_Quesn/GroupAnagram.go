package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	tmp := strings.Split(s, "")
	sort.Strings(tmp)
	return strings.Join(tmp, "")

}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := [][]string{}
	mp := map[string][]string{}

	for _, s := range strs {
		tmp := sortString(s)
		mp[tmp] = append(mp[tmp], s)

	}
	for _, x := range mp {
		result = append(result, x)
	}
	fmt.Println(result)

}

func groupAnagrams(strs []string) [][]string {
	sorted := map[string][]int{}
	res := [][]string{}
	for i, v := range strs {
		tmp := []byte(v)
		sort.Slice(tmp, func(a, b int) bool { return tmp[a] < tmp[b] })
		sorted[string(tmp)] = append(sorted[string(tmp)], i)
	}
	for _, v := range sorted {
		tmp := []string{}
		for _, index := range v {
			tmp = append(tmp, strs[index])
		}
		res = append(res, tmp)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	tmp := map[[26]int][]string{}
	for _, s := range strs {
		chars := [26]int{}
		for _, c := range s {
			chars[c-'a']++
		}
		tmp[chars] = append(tmp[chars], s)
	}
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}
