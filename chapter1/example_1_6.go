package main

import (
	"fmt"
	"time"
)

// func to exectute as go routine
func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count() // start go routine
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World!")
	time.Sleep(time.Millisecond * 5)
}
