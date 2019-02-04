package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(30 * time.Second) // create a channel that will receive a message when 30 seconds have elapsed
	echo := make(chan []byte)            // make a new channel for passing bytes from Stdin to Stdout. Because you haven't specified a size, this channel can hold only one messate at a time.
	go readStdin(echo)                   // statrs a goroutine to read Stdin, passes it our new channel for communicating
	for {                                // use a Select statement to pass data from Stdin to Stdout when received, or to shut down when the time-out event occurs
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) { // Takes a write-only channel (chan <-) and send any receivedinput to that channel
	for {
		data := make([]byte, 1024) // copies some data from stdin into data. Note that File.Read blocks until it receives data.
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data // sends the buffered data over the channel
		}
	}
}
