package main

import (
	"fmt"

	"github.com/sonnyb378/pointersChallenges/challenges/easy"
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

	fmt.Println("Increment Using Pointer")
	num := 5
	pointer := &num
	incrementValue := 1
	fmt.Printf("Counter: %v; incrementValue: %v\n", num, incrementValue)
	easy.Increment(pointer, incrementValue)

}
