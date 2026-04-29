package main

import (
	"bytes"
	"context"
	"fmt"
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

	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, err := conn.Read(buffer)
			if err != nil {
				panic(err)
			}

			if !bytes.Contains(buffer, CRLF) {
				message = append(message, buffer...)
			} else {
				fmt.Printf("Parse Line: %s\n", string(message))
				message = message[:0]
			}
		}

	}
}
