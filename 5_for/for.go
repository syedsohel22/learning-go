package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop

	for {
		fmt.Println("hello")
		break // to stop this
	}

	// classic for loop

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// break and continue is also work as expected

	// range
	fmt.Println("range")
	// it run till limit-1 here 10 is limit
	for i := range 10 {
		fmt.Println(i)
	}

}
