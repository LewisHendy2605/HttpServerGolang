package server

import (
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/LewisHendy2605/HttpServerGolang/internal/request"
	"github.com/LewisHendy2605/HttpServerGolang/internal/response"
)

type Server struct {
	closed  bool
	handler Handler
}

// Starts a new tcp server
func Serve(port uint16, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := &Server{closed: false, handler: handler}
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
	defer conn.Close()

	res := response.NewResponse()
	res.Headers.Set("Connection", "close")
	res.Headers.Set("Content-Length", "0")
	res.Headers.Set("Content-Type", "text/plain")

	req, err := request.RequestFromReader(conn)
	if err != nil {
		slog.Error("parsing request", "error", err)
		panic(err)
	}

	handlerErr := s.handler(res, req)
	if handlerErr != nil {
		slog.Error("handling request", "error", err)
		panic(handlerErr)
	}

	err = res.WriteStatusLine(conn, 200)
	if err != nil {
		slog.Error("writing status line", "error", err)
		panic(err)
	}

	err = res.WriteHeaders(conn)
	if err != nil {
		slog.Error("writing headers", "error", err)
		panic(err)
	}

}

// Closes server connection, and stops handling connection
func (s *Server) Close() error {
	s.closed = true
	return nil
}
