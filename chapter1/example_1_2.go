package main

import (
	"fmt"
)

// func with two named strings defined for return
func Names() (first string, second string) {
	return "Foo", "Bar" // two values returned by func
}

func main() {
	n1, n2 := Names()   // variable get two values
	fmt.Println(n1, n2) // print them.
}
