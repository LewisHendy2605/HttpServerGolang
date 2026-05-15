package server

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	client, conn := net.Pipe()
	defer client.Close()
	defer conn.Close()

	go func() {
		server := &Server{}
		server.handle(conn)
	}()

	buf := make([]byte, 4096)

	n, err := client.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "HTTP/1.1 200 OK\r\nconnection: close\r\ncontent-length: 0\r\n\r\n", string(buf[:n]))
}
