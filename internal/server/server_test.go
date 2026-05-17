package server

import (
	"io"
	"log/slog"
	"net"
	"testing"

	"github.com/LewisHendy2605/HttpServerGolang/internal/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	client, conn := net.Pipe()
	defer client.Close()

	var handler Handler = func(w io.Writer, req *request.Request) *HandlerError {
		slog.Info("http handler")

		return nil
	}

	go func() {
		server := &Server{handler: handler}
		server.handle(conn)
	}()

	// Write request to server
	_, err := client.Write([]byte(
		"GET / HTTP/1.1\r\n" +
			"Host: localhost\r\n" +
			"\r\n",
	))
	require.NoError(t, err)

	buf, err := io.ReadAll(client)
	require.NoError(t, err)

	assert.Equal(t, "HTTP/1.1 200 OK\r\nconnection: close\r\ncontent-length: 0\r\ncontent-type: text/plain\r\n\r\n", string(buf))
}
