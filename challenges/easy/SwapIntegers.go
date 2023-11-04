package easy

import (
	"fmt"
)

func SwapIntegers(a, b *int) {
	// Challenge 1: Swap Two Integers
	// Description: Write a Go function that swaps the values of two integers using pointers.

	// Input:
	// a := 5
	// b := 10

	// Output:
	// a is now 10
	// b is now 5

	// Incorrect Output:
	// a is still 5
	// b is still 10

	*a, *b = *b, *a

	fmt.Println()
	fmt.Println("After swap:")
	fmt.Printf("a is now %v\n", *a)
	fmt.Printf("b is now %v\n", *b)
}
