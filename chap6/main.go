package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func ex1() {
	fn := "Alex"
	ln := "Langevin"
	a := 35

	test1 := MakePerson(fn, ln, a)
	fmt.Println(test1)
	test2 := MakePersonPointer(fn, ln, a)
	fmt.Println(test2)
}

func UpdateSlice(sli []string, str string) {
	sli[len(sli)-1] = str
	fmt.Println("inside UpdateSlice: ", sli)
}

func GrowSlice(sli []string, str string) {
	sli = append(sli, str)
	fmt.Println("inside GrowSlice: ", sli)
}

func ex2() {
	sli := []string{"Hi there!", "please hire me for a dev role.", "I'm good at go!"}
	str := "I'm the best at go!"
	str2 := "Well, one of the best... ;)"

	fmt.Println("original slice: ", sli)
	UpdateSlice(sli, str)
	fmt.Println("after UpdateSlice: ", sli)

	GrowSlice(sli, str2)
	fmt.Println("after GrowSlice: ", sli)
}

func main() {
	ex1()
	// ./main.go:11:6: can inline MakePerson
	// ./main.go:19:6: can inline MakePersonPointer
	// ./main.go:38:6: can inline main
	// ./main.go:32:21: inlining call to MakePerson
	// ./main.go:33:13: inlining call to fmt.Println
	// ./main.go:34:28: inlining call to MakePersonPointer
	// ./main.go:35:13: inlining call to fmt.Println
	// ./main.go:11:17: leaking param: firstName to result ~r0 level=0
	// ./main.go:11:28: leaking param: lastName to result ~r0 level=0
	// ./main.go:19:24: leaking param: firstName
	// ./main.go:19:35: leaking param: lastName
	// ./main.go:20:9: &Person{...} escapes to heap
	// ./main.go:33:13: ... argument does not escape
	// ./main.go:33:14: test1 escapes to heap
	// ./main.go:34:28: &Person{...} escapes to heap
	// ./main.go:35:13: ... argument does not escape

	// The &Person{} returned from MakePersonPointer escapes to the heap. Any time a pointer is returned from a function, the pointer is returned on the stack, but the value it points to must be stored on the heap.
	//
	// Surprisingly, it also says that p escapes to the heap. The reason for this is that it is passed into fmt.Println. This is because the parameter to fmt.Println are ...any. The current Go compiler moves to the heap any value that is passed in to a function via a parameter that is of an interface type. (Interfaces are covered in Chapter 7.)

	ex2()
	// original slice:  [Hi there! please hire me for a dev role. I'm good at go!]
	// inside UpdateSlice:  [Hi there! please hire me for a dev role. I'm the best at go!]
	// after UpdateSlice:  [Hi there! please hire me for a dev role. I'm the best at go!]
	// inside GrowSlice:  [Hi there! please hire me for a dev role. I'm the best at go! Well, one of the best... ;)]
	// after GrowSlice:  [Hi there! please hire me for a dev role. I'm the best at go!]

	// both functions has the same backing array. When the array is passed as a parameter, the length and capacity values are also copied.
	// The Go runtime prevents the original slice from seeing those values since ther are beyond the length of the original slice.

}
