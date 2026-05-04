package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRequestLine(t *testing.T) {
	// Valid Request Line
	rl := &RequestLine{}
	_, err := rl.Parse([]byte("GET http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"))
	require.NoError(t, err)
	require.Equal(t, MethodGet, rl.Method)
	require.Equal(t, "http://www.w3.org/pub/WWW/TheProject.html", rl.RequestTarget)
	require.Equal(t, 1, rl.HttpVersion.Major)
	require.Equal(t, 1, rl.HttpVersion.Minor)

	// Valid Request Line
	rl = &RequestLine{}
	_, err = rl.Parse([]byte("GET https:/example.com HTTP/1.1"))
	require.NoError(t, err)
	require.Equal(t, MethodGet, rl.Method)
	require.Equal(t, "https:/example.com", rl.RequestTarget)
	require.Equal(t, 1, rl.HttpVersion.Major)
	require.Equal(t, 1, rl.HttpVersion.Minor)

	// Missing Method
	rl = &RequestLine{}
	_, err = rl.Parse([]byte("http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"))
	require.Error(t, err)

	// Missing request target
	rl = &RequestLine{}
	_, err = rl.Parse([]byte("GET  HTTP/1.1"))
	require.Error(t, err)

	// Missing version
	rl = &RequestLine{}
	_, err = rl.Parse([]byte("GET http://www.w3.org/pub/WWW/TheProject.html "))
	require.Error(t, err)

	// Invalid method
	rl = &RequestLine{}
	_, err = rl.Parse([]byte("GETT http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"))
	require.Error(t, err)
}
