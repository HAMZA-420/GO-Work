package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"a": "apple", "b": "banana"}

	fmt.Printf("a\t%v\n", a)

	var b = make(map[string]string)
	b["brand"] = "Ford"
	b["model"] = "Mustang"
	b["year"] = "1964"
	fmt.Printf(b["brand"])
}
