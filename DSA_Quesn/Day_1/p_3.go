// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

package main

import "fmt"

func solve(nums []int, target int) (int, int) {
	m := make(map[int]int, len(nums))
	for i, x := range nums {
		j, flag := m[x]
		if flag {
			return j, i
		}
		m[target-x] = i
	}
	return -1, -1

}

func main() {

	nums := []int{3, 2, 4}
	target := 6

	index1, index2 := solve(nums, target)

	fmt.Println(index1, index2)

}
