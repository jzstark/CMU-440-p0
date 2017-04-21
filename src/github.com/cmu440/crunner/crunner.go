package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	defaultHost = "localhost"
	defaultPort = 9999
)

// To test your server implementation, you might find it helpful to implement a
// simple 'client runner' program. The program could be very simple, as long as
// it is able to connect with and send messages to your server and is able to
// read and print out the server's response to standard output. Whether or
// not you add any code to this file will not affect your grade.
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", defaultHost+":"+strconv.Itoa(defaultPort))
	if err != nil {
		os.Exit(1)
	}

	//dial from local address to remote address
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("connect success")
	//send message
	getValue(conn)
	conn.Close()
}

func getValue(conn net.Conn) {
	words := "hello world\n"
	conn.Write([]byte(words))
	fmt.Print("send out: ", words)
}
