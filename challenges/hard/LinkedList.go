package hard

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Add(data int) {
	newNode := &Node{data, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) DeleteAt(index int) {

	if index < 0 || list.head == nil {
		fmt.Println("Error: Index out of range or List is empty")
		return
	}

	if index == 0 {
		current := list.head
		list.head = list.head.next
		fmt.Printf("Deleting position at index: %v with data %v\n", index, current.data)
		return
	}

	i := 0
	value := 0
	current := list.head
	previous := current

	for current != nil && i < index {
		previous = current
		current = current.next
		i++
	}

	if current != nil {
		previous.next = current.next
		value = current.data
	}

	fmt.Printf("Deleting position at index: %v with data %v\n", index, value)
}

func (list *LinkedList) Search(data int) {

	index := 0

	if list.head.data == data {
		fmt.Printf("> %v found in list at index %v\n", data, index)
		return
	}
	current := list.head

	for current != nil {
		if current.data == data {
			fmt.Printf(">> %v found in list at index %v\n", data, index)
			return
		}
		current = current.next
		index++
	}
	fmt.Printf(">> %v not found!\n", data)
}

func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
	fmt.Println()
}

func CreateLinkedList() {

	// Challenge 8: Implement a Linked List with Pointers
	// Description: Implement a singly linked list from scratch in Go with functions to add, delete, and search for elements using pointers.

	// **Input and Output: Varies based on operations performed on the linked list.

	// Incorrect Output:
	// The linked list is not properly implemented.

	myList := LinkedList{}
	myList.Add(1)
	myList.Add(2)
	myList.Add(3)
	myList.Add(4)
	myList.Add(5)

	myList.Print()

	fmt.Println()

}
