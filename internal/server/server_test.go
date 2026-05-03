package server

import (
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	client, conn := net.Pipe()

	ctx := t.Context()

	go func() {
		defer client.Close()
		fmt.Fprintf(client, "GET /coffee HTTP/1.1\r\n")
	}()

	for line := range ReadLines(conn, ctx) {
		fmt.Printf("Request Line: %s\n", line)
	}
}
