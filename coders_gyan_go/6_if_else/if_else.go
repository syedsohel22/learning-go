package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 15 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a child")
	}

	// in go we dont use () in loop and if else

	// Logical operators
	// || OR operator
	// && AND operator
	// ! NOT operator

	// we can declare a variable in if construct

	// go does not have ternary as of go 1.22
}
