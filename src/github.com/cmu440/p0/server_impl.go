// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type keyValueServer struct {
	// TODO: implement this!
	listener net.Listener
	port     int16
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	log.Print("New a kyeValueServer")
	return &keyValueServer{
		listener: nil,
		port:     0,
	}
}

func (kvs *keyValueServer) Start(port int) error {
	// TODO: implement this!
	var err error
	kvs.listener, err = net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	checkError(err)

	for {
		conn, err := kvs.listener.Accept()
		if err != nil {
			continue
		}

		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)

		if err != nil {
			return nil
		}

		fmt.Printf("%d %s\n", n, buffer)
	}

	return nil
}

func (kvs *keyValueServer) Close() {
	// TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
	// TODO: implement this!
	return -1
}

// TODO: add additional methods/functions below!

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
