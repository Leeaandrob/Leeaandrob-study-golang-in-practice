// this example has problem with race condition
// In a race condition, two things are

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup // you'll use a wait group to monitor a gorup of go routines
	w := newWords()
	for _, f := range os.Args[:1] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:") // at the end the program will print what you found.
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

type words struct { // you track words in a struct. You could use a map type, but using a struct here makes	the next refactor easie
	found map[string]int
}

func newWords() *words { // create  a new words easier
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// open a file, parse its contents, and count the words that appear.
// copy function does all the copying for you.
// Scanner is a useful tool for passing files like this.
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
