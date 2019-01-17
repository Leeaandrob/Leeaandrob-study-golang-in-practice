package main

import (
	"fmt"
	"time"
)

// func with an int type channed passed
func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c // waits to value to come in
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) // a channel is created
	a := []int{8, 7, 5, 3, 9, -1}

	go printCount(c) // start the go routine

	for _, v := range a {
		c <- v // pass ints into channel
	}

	time.Sleep(time.Millisecond * 1) // main pauses before ending
	fmt.Println("End of main")
}
