package main

import "fmt"

func main() {
	const name string = "sohel"

	// cannot assign to name (neither addressable nor a map index expression)compilerUnassignableOperand
	// name = "javascript"

	// you can't you shorthand := out side function

	// constant grouping

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(name, port, host)
}
