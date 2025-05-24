// You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the
//ith position moves to indices[i] in the shuffled string.

// Return the shuffled string.

// Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcdoe"
// Explanation: As shown, "codeleet" becomes "leetcdoe" after shuffling.

// Input: s = "abc", indices = [0,1,2]
// Output: "abc"
// Explanation: After shuffling, each character remains in its position.

package main

import "fmt"

func main() {
	str := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	for i := 0; i < len(indices); i++ {
		fmt.Printf("%c", str[indices[i]])

	}
}
