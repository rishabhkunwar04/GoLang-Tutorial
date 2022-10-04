// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original
//letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]
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
