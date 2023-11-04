package main

import (
	"fmt"

	// "github.com/sonnyb378/pointersChallenges/challenges/easy"
	"github.com/sonnyb378/pointersChallenges/challenges/intermediate"
)

func main() {

	// EASY
	// fmt.Println("Challenge 1: Swap Two Integers")
	// a := 5
	// b := 10
	// fmt.Println("Before swap:")
	// fmt.Printf("a is %v\n", a)
	// fmt.Printf("b is %v\n", b)
	// easy.SwapIntegers(&a, &b)

	// fmt.Println("Modify Slice Element")
	// slice := []int{1, 2, 3, 4, 5}
	// position := 2
	// newValue := 10
	// fmt.Println("Before: ", slice, "; position is ", position, "; newValue is ", newValue)
	// easy.ModifySlice(slice, position, newValue)

	// fmt.Println("Increment Using Pointer")
	// num := 5
	// pointer := &num
	// incrementValue := 1
	// fmt.Printf("Counter: %v; incrementValue: %v\n", num, incrementValue)
	// easy.Increment(pointer, incrementValue)

	// INTERMEDIATE
	// fmt.Println("Challenge 4: Swap Two Strings")
	// str1 := "Hello"
	// str2 := "World"
	// fmt.Printf("Before swap: str1 is %v; str2 is %v\n", str1, str2)
	// intermediate.SwapStrings(&str1, &str2)

	// fmt.Println("Challenge 5: Update struct fields with pointers")
	// name := "Alice"
	// age := 25
	// person := &intermediate.Person{Name: name, Age: age}
	// fmt.Println()
	// fmt.Println("Before update: ", *person)
	// intermediate.UpdateStruct(person, "Bob", 30)
	// fmt.Println("After update: ", *person)
	// fmt.Println()

	fmt.Println("Challenge 6: Modify Map Value Using Pointer")
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	key := "b"
	newValue := 10
	fmt.Printf("Before --> %p:%v\n", &myMap, myMap)
	intermediate.ModifyMapValue(&myMap, key, newValue)
	fmt.Printf("After --> %p:%v\n", &myMap, myMap)
}
