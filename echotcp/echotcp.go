/*
	Written while following along with the talk "Go: code that grows with grace". Check out the talk
	here: https://vimeo.com/53221560
*/

package main

import (
	"io"
	"log"
	"net"
)

const listenerAddr = "localhost:4000"

func main() {
	listener, err := net.Listen("tcp", listenerAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go io.Copy(connection, connection)
	}
}
