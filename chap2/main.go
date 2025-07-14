package main

import (
	"fmt"
	"math"
)

func ex2() {
	fmt.Println("--- exercise 2 ---")
	i := 20
	f := float64(i)

	fmt.Println("i: ", i)
	fmt.Println("f: ", f)
	fmt.Println()
}

func ex3() {
	fmt.Println("--- exercise 3 ---")
	// is there arithmetic overflow in go?
	// Let's see!

	// Don't write the by hand.
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	fmt.Println("b: ", b+1)
	fmt.Println("smallI: ", smallI+1)
	fmt.Println("bigI: ", bigI+1)

	// There sure is!
	// Hurray!!!
}

func main() {
	ex2()
	ex3()
}
