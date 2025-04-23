package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("/---- ex1 ----/")
	ex1()

	fmt.Println()
	fmt.Println("/---- ex2 ----/")
	ex2()

	fmt.Println()
	fmt.Println("/---- ex3 ----/")
	ex3()
}

func ex1() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Printf("i: %d\nf: %f\n", i, f)
}

func ex2() {
	const value = 55
	var i int = value
	var f float64 = value

	fmt.Println("i: ", i)
	fmt.Println("f: ", f)
}

func ex3() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	fmt.Println("b: ", b+1)
	fmt.Println("smallI: ", smallI+1)
	fmt.Println("bigI: ", bigI+1)
}
