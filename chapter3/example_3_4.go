package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup // A waitgroup doesn't need to be initialized
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1) // For every file you add, you tell the wait group that you’re waiting for one more compress operation.
		go func(filename string) {
			compress(filename) // this function calls compress and then notifies the
			wg.Done()          // wait group that it's done.
		}(file) // because you're calling a goroutine in a for loop, you need to do a little trickery with the parameter passing.
	}
	wg.Wait() // the outer goroutine (main) waits untill all the compressing goroutines have called wg.Done
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
