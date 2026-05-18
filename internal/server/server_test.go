package server

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/LewisHendy2605/HttpServerGolang/internal/request"
	"github.com/LewisHendy2605/HttpServerGolang/internal/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	client, conn := net.Pipe()
	defer client.Close()

	timeString := time.Now().Format(time.RFC3339)

	var handler Handler = func(res *response.Response, req *request.Request) *HandlerError {
		res.Headers.Set("Date", timeString)

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

	t.Log(string(buf))

	assert.Equal(t, fmt.Sprintf("HTTP/1.1 200 OK\r\nconnection: close\r\ncontent-length: 0\r\ncontent-type: text/plain\r\ndate: %s\r\n\r\n", timeString), string(buf))
}
