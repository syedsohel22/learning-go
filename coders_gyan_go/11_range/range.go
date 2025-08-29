package main

import "fmt"

// interating over data structures
func main() {

	nums := []int{6, 7, 8}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	sum := 0
	for i, num := range nums {
		fmt.Println(num, i)
		sum = sum + num
	}

	fmt.Println(sum)

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// range on string data
	for i, c := range "golang" {
		// here c will return unicode code
		// point rune
		// i is starting point (byte) of rune
		// fmt.Println(i,c)
		fmt.Println(i, c, string(c))
	}
}
