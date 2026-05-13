package server

import (
	"log"
	"net"

	"github.com/LewisHendy2605/HttpServerGolang/internal/request"
)

type Server struct {
}

// Starts tcp server
func ListenAndServe(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error accepting connection: %v", err)
		}

		_, err = request.RequestFromReader(conn)

		// TODO: What to do here, do we return 400, 500??
		// the server shouldn't stop if the request is bad
		// but also we dnt want to leak internal error messages
		if err != nil {
			panic(err)
		}
	}
}
