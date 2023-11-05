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

	list.Reverse()

	fmt.Println()
	fmt.Println("Reversed: ")
	list.Print()

}

func (list *LinkedList) Reverse() {
	var previousHolder *Node = nil
	current := list.head
	var nextHolder *Node = nil

	list.Print()

	for current != nil {
		nextHolder = current.next
		current.next = previousHolder
		previousHolder = current
		current = nextHolder
	}

	list.head = previousHolder
}
