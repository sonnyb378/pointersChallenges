package intermediate

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: \"%s\", Age: %d}", p.Name, p.Age)
}

func UpdateStruct(person *Person, name string, age int) {
	// Challenge 5: Update Struct Fields with Pointers
	// Description: Write a function that takes a pointer to a struct and updates one or more of its fields.

	// Input:
	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	// initial struct values
	// Name: Alice
	// Age: 25

	// UpdatePerson(person, "Bob", 30)

	// Output:
	// Person{Name: "Bob", Age: 30}

	// Incorrect Output:
	// Person{Name: "Alice", Age: 25}

	person.Name = name
	person.Age = age

}
