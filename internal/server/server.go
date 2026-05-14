package server

import (
	"fmt"
	"io"
	"net"

	"github.com/LewisHendy2605/HttpServerGolang/internal/response"
)

type Server struct {
	closed bool
}

// Starts a new tcp server
func Serve(port uint16) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := &Server{closed: false}
	go s.runServer(listener)

	return s, nil
}

// Accepts new tcp connection and passes to go routine to handle
func (s *Server) runServer(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil || s.closed {
			return
		}

		go s.handle(conn)
	}
}

// Handles a new tcp connection
func (s *Server) handle(conn io.ReadWriteCloser) {
	res := response.NewResponse()
	res.Headers.Set("Content-Length", "0")
	res.Ok()

	conn.Write(res.Byte())
}

// Closes server connection, and stops handling connection
func (s *Server) Close() error {
	s.closed = true
	return nil
}
