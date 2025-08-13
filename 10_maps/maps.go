package main

import "fmt"

func main() {
	// Maps
	// var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	// b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	// Maps with make function
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Println(a)
	fmt.Println(b)
}
