package main

import "fmt"

func main(){
	// declaration of an array

	var a [5]int
	a[0]=1
	a[1]=2
	a[2]=3
	a[3]=4
	a[4]=5

	// length of an array
	fmt.Print(len(a))
	// accessing and adding element the array
	fmt.Println(a[0])

	// 2d array

	var b [2][3]int
	b[0][0]=1
	b[0][1]=2
	b[0][2]=3
	b[1][0]=4
	b[1][1]=5
	b[1][2]=6

	fmt.Println(b)

	declaration := [5]int{1,2,3,4,5}
	fmt.Println(declaration)

}
