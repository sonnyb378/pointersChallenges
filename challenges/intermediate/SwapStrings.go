package intermediate

import (
	"fmt"
)

func SwapStrings(str1, str2 *string) {

	// Challenge 4: Swap Two Strings
	// Description: Write a Go function that swaps the values of two strings using pointers.

	// Input:
	// str1 := "Hello"
	// str2 := "World"

	// Output:
	// str1 is now "World"
	// str2 is now "Hello"

	// Incorrect Output:
	// str1 is still "Hello"
	// str2 is still "World"

	*str1, *str2 = *str2, *str1

	fmt.Printf("str1 is now %v\n", *str1)
	fmt.Printf("str2 is now %v\n", *str2)
}
