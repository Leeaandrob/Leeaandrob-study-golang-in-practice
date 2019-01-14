package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "golang.org:80")         // connects over ctp
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")         // sends strings over the connection
	status, _ := bufio.NewReader(conn).ReadString('\n') // Prints the first response line
	fmt.Println(status)
}
