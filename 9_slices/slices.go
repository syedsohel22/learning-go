package main

import "fmt"

func main() {
	myslice := []int{}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice1 := []int{1, 2, 3}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

	//len() function - returns the length of the slice (the number of elements in the slice)
	//cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	sliceExample()

	sliceWithMake()
}

func sliceExample() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

func sliceWithMake() {
	myslice := make([]int, 5, 10)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)
}
