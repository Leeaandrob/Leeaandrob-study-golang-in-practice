package main

import (
	"fmt"
)

// func with two strings defined for return
func Names() (string, string) {
	return "Foo", "Bar" // two values returned by func
}

func main() {
	n1, n2 := Names()   // variable get two values
	fmt.Println(n1, n2) // print them.

	n3, _ := Names() // get first value and print them.
	fmt.Println(n3)
}
