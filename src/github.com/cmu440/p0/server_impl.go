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
	cliPool  []*net.Conn
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	log.Print("New a kyeValueServer")
	return &keyValueServer{
		listener: nil,
		port:     0,
		cliPool:  make([]*net.Conn),
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

		go connHandler(conn)

	}

	return nil
}

func (kvs *keyValueServer) Close() {
	// TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
	// TODO: implement this!
	return len(kvs.cliPool)
	//	return -1
}

// TODO: add additional methods/functions below!

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func connHandler(conn net.Conn) {
	fmt.Printf("start a goroutine to handle: %s\n", conn.RemoteAddr().String())
	buff := make([]byte, 2048)

	_, err := conn.Read(buff)

	if err != nil {
		return
	}

	fmt.Print("received: ", string(buff))
}
