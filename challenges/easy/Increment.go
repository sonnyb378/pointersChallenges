package easy

import (
	"fmt"
)

func Increment(pointer *int, incrementValue int) {

	// Challenge 3: Increment Using Pointer
	// Description: Write a Go function that takes a pointer to an integer and increments the value it points to.

	// Input:
	// num := 5
	// pointer := &num

	// Output:
	// num is now 6

	// Incorrect Output:
	// num is still 5

	*pointer = *pointer + incrementValue

	fmt.Printf("counter is now %v\n", *pointer)
}
