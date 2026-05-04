package request

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRequest(t *testing.T) {
	// Test good request
	reader := strings.NewReader("GET / HTTP/1.1\r\nHost: example.com\r\nUser-Agent: golang_application\r\n\r\nHello")
	req, err := RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, req)
	require.Equal(t, "GET", req.RequestLine.Method)
	require.Equal(t, "/", req.RequestLine.RequestTarget)
	require.Equal(t, 1, req.RequestLine.HttpVersion.Major)
	require.Equal(t, 1, req.RequestLine.HttpVersion.Minor)

	val, ok := req.Headers.Get("HOST")
	require.True(t, ok)
	require.Equal(t, "example.com", val)

	val, ok = req.Headers.Get("user-agent")
	require.True(t, ok)
	require.Equal(t, "golang_application", val)
}
