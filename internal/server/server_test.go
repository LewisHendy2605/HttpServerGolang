package server

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	client, conn := net.Pipe()
	defer client.Close()

	go func() {
		server := &Server{}
		server.handle(conn)
	}()

	buf, err := io.ReadAll(client)
	require.NoError(t, err)

	assert.Equal(t, "HTTP/1.1 200 OK\r\nconnection: close\r\ncontent-length: 0\r\ncontent-type: text/plain\r\n\r\n", string(buf))
}
