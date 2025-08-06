package main

import (
	"fmt"
)

// Function that modifies the first element of a slice
func modifySlice(s []int) {
	s[0] = 100 // Modify the zeroth element
	fmt.Println("Inside function, slice:", s)
}
func main() {
	arr := [3]int{1, 2, 3}
	s := arr[:] // Create a slice referencing the whole array

	fmt.Println("Before function call, array:", arr)
	modifySlice(s)
	fmt.Println("After function call, array:", arr)
}

/*
Before function call, array: [1 2 3]
Inside function, slice: [100 2 3]
After function call, array: [100 2 3]

*/
