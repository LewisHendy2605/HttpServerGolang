package main

import (
	"bytes"
	"context"
	"net"
)

var CRLF = []byte("\r\n")

func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go HandleConn(conn, context.Background())
	}
}

func HandleConn(conn net.Conn, ctx context.Context) {
	buffer := make([]byte, 2)
	message := make([]byte, 1057)
	numBytes := 0

	for {
		select {
		case <-ctx.Done():
			return
		default:
			bytes_read, err := conn.Read(buffer)
			if err != nil {
				panic(err)
			}

			if !bytes.Contains(buffer, CRLF) {
				message = append(message, buffer...)
				numBytes += bytes_read
			} else {
				ParseRequestLine(message)
				message = message[:0]
			}
		}

	}
}
