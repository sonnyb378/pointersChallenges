package intermediate

import (
	"fmt"
)

func ModifyMapValue(m *map[string]int, key string, newValue int) {

	// Challenge 6: Modify Map Value Using Pointer
	// Description: Write a Go function that takes a pointer to a map and modifies the value associated with a given key.

	// Input:
	// myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	// key := "b"
	// newValue := 10

	// Output:
	// myMap is now map[a:1 b:10 c:3]

	// Incorrect Output:
	// myMap is still map[a:1 b:2 c:3]

	(*m)[key] = newValue

	fmt.Printf("Modify --> %p:%v\n", m, *m)

}
