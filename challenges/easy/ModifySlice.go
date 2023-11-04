package easy

import (
	"fmt"
)

func ModifySlice(arr []int, position int, newValue int) {

	// Challenge 2: Modify Slice Element
	// Description: Write a function that takes a slice and a position as input, and changes the value of the element at that position to a new value.

	// Input:
	// slice := []int{1, 2, 3, 4, 5}
	// position := 2
	// newValue := 10

	// Output:
	// slice is now [1 2 10 4 5]

	// Incorrect Output:
	// slice is still [1 2 3 4 5]

	if position >= 0 && position < len(arr) {
		arr[position] = newValue
	}

	fmt.Printf("Slice is now %v\n", arr)
}
