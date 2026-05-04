package server

import (
	"log"
	"net"
)

// Starts tcp server
func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		_, err = listener.Accept()
		if err != nil {
			log.Fatalf("error accepting connection: %v", err)
		}
	}
}
