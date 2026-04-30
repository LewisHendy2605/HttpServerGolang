package main

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	client, conn := net.Pipe()

	ctx, cancel := context.WithCancel(context.Background())

	go HandleConn(conn, ctx)
	defer conn.Close()

	fmt.Fprintf(client, "GET / HTTP/1.0\r\n")

	cancel()
}
