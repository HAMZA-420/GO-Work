package main

import (
	"fmt"
)

func main() {
	// var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	// val1, ok1 := a["brand"] // Checking for existing key and its value
	// val2, ok2 := a["color"] // Checking for non-existing key and its value
	// val3, ok3 := a["day"]   // Checking for existing key and its value
	// _, ok4 := a["model"]    // Only checking for existing key and not its value

	// fmt.Println(val1, ok1)
	// fmt.Println(val2, ok2)
	// fmt.Println(val3, ok3)
	// fmt.Println(ok4)

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	//loop with no order
	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	var b []string //defining the order
	b = append(b, "one", "two", "three", "four")
	for _, element := range b {
		fmt.Printf("%v : %v, ", element, a[element])
	}
}
