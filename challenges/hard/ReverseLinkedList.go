package hard

import (
	"fmt"
)

func ReverseLinkedList(list *LinkedList) {

	// Challenge 7: Reverse a Linked List Using Pointers
	// Description: Write a Go function that reverses a singly linked list using pointers.

	// Input:
	// 1 -> 2 -> 3 -> 4 -> 5

	// Output:
	// 5 -> 4 -> 3 -> 2 -> 1

	// Incorrect Output:
	// 1 -> 2 -> 3 -> 4 -> 5

	list.Print()
	fmt.Printf("%p\n", list)

}