package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	// A basic message
	// msg := "Here is a message"

	// A message with command line arguments
	msgP := flag.String("msg", "Hello", "Message to send")
	flag.Parse()

	// Dial connects to the stated address on the stated protocol.
	// It follows the form:
	// func Dial(network string, address string) (Conn, error)
	// Over here, the error, if any is returned, is ignored using _

	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	// Fprintln is like Println, but you can choose the destination of the output.
	// For example, Fprintln(os.Stdout, "Hello") prints to the console.
	// Over here, what we are doing is sending our message to the connection, rather
	// than the console.

	fmt.Fprintln(conn, *msgP)
}
