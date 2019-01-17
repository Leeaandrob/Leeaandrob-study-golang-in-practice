package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://example.com/") // make an http get request

	body, _ := ioutil.ReadAll(resp.Body) // reads the body from the response

	fmt.Println(string(body))
	resp.Body.Close() // close the connection
}
