package main

import (
	"time"
)

func main() {
	time.Sleep(5 * time.Second)          // block for five seconds
	sleep := time.After(5 * time.Second) // creates a channel that will get notified in five seconds, then block until that channel receives a notification
	<-sleep
}
