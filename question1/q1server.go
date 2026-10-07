package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	// func Listen(network string, address string) (Listener, error)
	// In essence, when we say net.Listen("tcp", ":8030"), we are saying that we
	// are using the TCP/IP protocol on port 8030.
	// What is returned is a Listener (which is denoted ln over here) and an error, if there
	// is one which is returned. Here, it is just ignored by replacing its output with a _

	ln, _ := net.Listen("tcp", ":8030")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
