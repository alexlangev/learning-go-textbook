package main

import (
	"fmt"
)

type Person struct {
	FirstName	string
	LastName	string
	Age			int
}
	
func MakePerson (firstName string, lastName string, age int) Person {
	person := Person {
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}

	return person
}

func MakePersonPointer (firstName string, lastName string, age int) *Person {
	person := Person {
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}

	return &person
}

func main() {
	me := MakePerson("Alex", "Langevin", 34)
	mePointer := MakePersonPointer("Sara", "Langevin", 38)

	fmt.Println(me)
	fmt.Println(mePointer)
}
