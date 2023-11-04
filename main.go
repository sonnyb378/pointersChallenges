package main

import (
	"fmt"

	"github.com/sonnyb378/pointersChallenges/challenges/easy"
)

func main() {

	// EASY
	a := 5
	b := 10
	fmt.Println("Before swap:")
	fmt.Printf("a is %v\n", a)
	fmt.Printf("b is %v\n", b)
	easy.SwapIntegers(&a, &b)

}
